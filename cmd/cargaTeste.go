package cmd

import (
	"fmt"
	"net/url"

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
	fmt.Printf("URL: %s\n", parsedUrl)
	fmt.Printf("Concorrência: %d\n", concurrency)
	fmt.Printf("Requisições %d\n", requests)
}
