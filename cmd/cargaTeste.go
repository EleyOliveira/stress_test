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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: TesteCarga,
}

func init() {
	rootCmd.AddCommand(cargaTesteCmd)
	cargaTesteCmd.Flags().IntP("requests", "r", 1, "Quantidade de requisições")

	cargaTesteCmd.Flags().StringP("url", "u", "", "URL para realizar as requisições")
	cargaTesteCmd.MarkFlagRequired("url")

	cargaTesteCmd.Flags().IntP("concurrency", "c", 1, "Quantidade de requisições concorrentes")

	//cargaTesteCmd.MarkFlagRequired("requests")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cargaTesteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cargaTesteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func TesteCarga(cmd *cobra.Command, args []string) {
	fmt.Println("cargaTeste called")
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
	fmt.Printf("URL: %s\n", parsedUrl)
	fmt.Printf("Concorrência: %d\n", concurrency)
	fmt.Printf("Requisições %d\n", requests)
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
	formatarRelatorio(statusCodes, tempoDecorrido)
}

func formatarRelatorio(statusCodes map[int]int, tempoDecorrido time.Duration) {
	fmt.Println("Tempo decorrido: ", tempoDecorrido)

	for status, quantidade := range statusCodes {
		if status == -1 {
			fmt.Printf("Erros: %d\n", quantidade)
		} else {
			fmt.Printf("Status %d: %d\n", status, quantidade)
		}
	}
}
