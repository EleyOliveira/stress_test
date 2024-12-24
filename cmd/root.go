package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "my-cli-project",
    Short: "A brief description of your application",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of your application.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to my CLI application!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // Define flags and configuration settings here.
    rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.my-cli-project.yaml)")
}