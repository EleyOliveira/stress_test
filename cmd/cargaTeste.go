package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// cargaTesteCmd represents the cargaTeste command
var cargaTesteCmd = &cobra.Command{
	Use:   "cargateste",
	Short: "Esse comando realiza um teste de carga",
	Long: `Esse comando permite que seja informada uma URL onde serão enviadas as requisições, o total de requisições
	e a quantidade de requisições concorrentes.
	Por exemplo: cargateste -u http://localhost:8020 -t 100 -c 10`,
	Run: TesteCarga,
}

func init() {
	rootCmd.AddCommand(cargaTesteCmd)
	cargaTesteCmd.Flags().IntP("requests", "r", 1, "Quantidade de requisições")

	cargaTesteCmd.Flags().StringP("url", "u", "", "URL para realizar as requisições")
	cargaTesteCmd.MarkFlagRequired("url")

	cargaTesteCmd.Flags().IntP("concurrency", "c", 1, "Quantidade de requisições concorrentes")
}

func TesteCarga(cmd *cobra.Command, args []string) {

	requests, err := cmd.Flags().GetInt("requests")
	if err != nil {
		fmt.Println("Erro ao obter a quantidade de requisições:", err)
		return
	}
	if requests <= 0 {
		fmt.Println("A quantidade de requisições deve ser maior que zero")
		return
	}
	urlInformada, err := cmd.Flags().GetString("url")
	if err != nil {
		fmt.Println("Erro ao obter a URL:", err)
		return
	}
	parsedUrl, err := url.ParseRequestURI(urlInformada)
	if err != nil {
		fmt.Println("Erro ao analisar a URL:", err)
		return
	}
	concurrency, err := cmd.Flags().GetInt("concurrency")
	if err != nil {
		fmt.Println("Erro ao obter a quantidade de requisições concorrentes:", err)
		return
	}
	if concurrency <= 0 {
		fmt.Println("A quantidade de requisições concorrentes deve ser maior que zero")
		return
	}

	criarRequisicoes(parsedUrl, requests, concurrency)
}

func criarRequisicoes(url *url.URL, requests int, concurrency int) {
	startTime := time.Now()

	var wg sync.WaitGroup
	chlimite := make(chan struct{}, concurrency)

	statusCodes := make(map[int]int)
	var mutex sync.Mutex

	for i := 0; i < requests; i++ {
		wg.Add(1)
		chlimite <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-chlimite }()

			resp, err := http.Get(url.String())
			if err != nil {
				fmt.Println("Erro ao fazer a requisição:", err)

				mutex.Lock()
				statusCodes[-1]++
				mutex.Unlock()

				return
			}
			defer resp.Body.Close()

			mutex.Lock()
			statusCodes[resp.StatusCode]++
			mutex.Unlock()

		}()
	}
	wg.Wait()
	tempoDecorrido := time.Since(startTime)
	formatarRelatorio(statusCodes, tempoDecorrido, url, concurrency, requests)
}

func formatarRelatorio(statusCodes map[int]int, tempoDecorrido time.Duration, parsedUrl *url.URL, concurrency int, requests int) {

	fmt.Printf("URL: %s\n", parsedUrl)
	fmt.Printf("Concorrência: %d\n", concurrency)
	fmt.Printf("Requisições: %d\n", requests)
	fmt.Println("Tempo decorrido: ", tempoDecorrido)

	for status, quantidade := range statusCodes {
		if status == -1 {
			fmt.Printf("Quantidade de erros: %d\n", quantidade)
		} else {
			fmt.Printf("Quantidade de requisições com StatusCode %d: %d\n", status, quantidade)
		}
	}
}
