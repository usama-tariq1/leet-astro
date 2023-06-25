/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/usama-tariq1/leet-astro/helper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve leet-gin Application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var console = helper.Console{}

		console.Log("Info", "Starting Engine")

		path := helper.GetWD()

		// Execute "go run main.go" command in the working directory
		runCmd := exec.Command("go", "run", "main.go")
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Dir = path
		err := runCmd.Run()
		if err != nil {
			fmt.Println("Failed to run 'go run main.go' command:", err)
			return
		}

		console.Log("Success", "Application started")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
