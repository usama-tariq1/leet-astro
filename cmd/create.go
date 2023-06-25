package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/usama-tariq1/leet-astro/artisan"
	"github.com/usama-tariq1/leet-astro/helper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create controller or model scaffolding",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var fileType string
		var name string

		if len(args) < 2 {
			fileType = args[0]
		} else if len(args) >= 2 {
			fileType = args[0]
			name = args[1]
		} else {
			console.Log("Error", "Create requires two arguments: Type and Name")
			console.Log("Error", "Example: leet-astro create controller UserController")
			return
		}

		fileTypes := []string{"controller", "model", "router"}

		if helper.Contains(fileType, fileTypes) {
			if fileType == "controller" {
				modelName, _ := cmd.Flags().GetString("model")
				if modelName == "" {
					artisan.CreateController(name)
				} else {
					artisan.CreateResourceController(name, modelName)
				}
			} else if fileType == "model" {
				artisan.CreateModel(name)
			} else if fileType == "router" {
				artisan.CreateRouter(name)
			}

		} else {
			console.Log("Error", "Unknown command for create")
			console.Log("Info", "Possible terms are:")
			console.Log("Info", strings.Join(fileTypes, ", "))
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().String("model", "", "Model name for creating a controller")
}
