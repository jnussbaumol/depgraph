/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// walkCmd represents the walk command
var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("walk called")
		rawPath := "./src/main/java"
		if len(args) > 1 {
			fmt.Println("Error: too many args provided.")
			return
		}
		if len(args) == 1 {
			rawPath = args[0]
		}
		path, err := filepath.Abs(rawPath)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(walkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
