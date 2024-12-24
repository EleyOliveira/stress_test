/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cargaTesteCmd represents the cargaTeste command
var cargaTesteCmd = &cobra.Command{
	Use:   "cargaTeste",
	Short: "Esse comando realiza um teste de carga",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Printf("Realizando %d requisições\n", requests)
	},
}

func init() {
	rootCmd.AddCommand(cargaTesteCmd)
	cargaTesteCmd.Flags().IntP("requests", "r", 1, "Quantidade de requisições")
	cargaTesteCmd.MarkFlagRequired("requests")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cargaTesteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cargaTesteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
