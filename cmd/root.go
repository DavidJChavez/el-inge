package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "elinge help",
	Long:  "elinge help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Script para instalar todo lo necesario para chambear")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
