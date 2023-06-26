/*
Copyright © 2022 NAME HERE <usama.tariq1337@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [ProjectName] [Mod]",
	Short: "Create project scaffolding",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		mod := "github.com/user"
		if len(args) > 1 {
			mod = args[1]
		}
		// Create a new folder with the given project name
		err := os.Mkdir(projectName, 0755)
		if err != nil {
			fmt.Println("Failed to create project folder:", err)
			return
		}

		// Change the current working directory to the project folder
		err = os.Chdir(projectName)
		if err != nil {
			fmt.Println("Failed to change directory to project folder:", err)
			return
		}

		// Run the command to pull the GitHub branch
		githubProject := "https://github.com/usama-tariq1/leet-gin.git"
		clonecmd := exec.Command("git", "clone", "-b", "main", githubProject, ".")
		clonecmd.Stdout = os.Stdout
		clonecmd.Stderr = os.Stderr
		err = clonecmd.Run()
		if err != nil {
			fmt.Println("Failed to pull GitHub branch:", err)
			return
		}

		// Modify the module name in go.mod
		moduleName := fmt.Sprintf("%s/%s", mod, projectName)
		modCmd := exec.Command("go", "mod", "edit", "-module", moduleName)
		modCmd.Stdout = os.Stdout
		modCmd.Stderr = os.Stderr
		err = modCmd.Run()
		if err != nil {
			fmt.Println("Failed to modify module name:", err)
			return
		}

		// Remove the .git folder
		gitPath := filepath.Join(projectName, ".git")
		err = os.RemoveAll(gitPath)
		if err != nil {
			fmt.Println("Failed to remove .git folder:", err)
			return
		}

		// Print success message
		fmt.Println("Project scaffolding created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
