package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var greeting string

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			if greeting != "" {
				fmt.Println(greeting)
			} else {
				fmt.Println(viper.GetString("greeting"))
			}
		},
	}

	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting message")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
