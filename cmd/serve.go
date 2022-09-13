/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

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

		path := helper.GetWD()

		// f := utils.DocValue(path)
		// router := gin.Default()

		// s := &http.Server{
		// 	Addr:           ":8080",
		// 	Handler:        router,
		// 	ReadTimeout:    10 * time.Second,
		// 	WriteTimeout:   10 * time.Second,
		// 	MaxHeaderBytes: 1 << 20,
		// }
		// s.ListenAndServe()
		err := http.ListenAndServe("localhost:8080", http.FileServer(http.Dir(path)))
		if err != nil {
			fmt.Println(err.Error())
		}
		console.Log("Success", "Router started")
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
