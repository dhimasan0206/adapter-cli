/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/spf13/cobra"
)

var moduleName, author, email, out string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		data := map[string]string{
			"moduleName": moduleName,
			"author":     author,
			"email":      email,
		}

		templates := "internal/templates"
		filepath.Walk(templates, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fmt.Println(out)
			target := strings.Replace(strings.Replace(path, ".mustache", "", 1), "internal/templates", out, 1)
			fmt.Println(target)

			if info.IsDir() {
				fmt.Printf("Directory: %s\n", path)
				err := os.MkdirAll(target, 0777)
				if err != nil {
					fmt.Println("Error creating directory:", err)
					return err
				}
				fmt.Println(target + " directory has been generated successfully.")
				return nil
			}

			fmt.Printf("File: %s\n", path)
			result, err := mustache.RenderFile(path, data)
			if err != nil {
				fmt.Println("Error rendering template:", err)
				return err
			}
			err = os.WriteFile(target, []byte(result), 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return err
			}
			fmt.Println(target + " has been generated successfully.")
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Module name")
	generateCmd.Flags().StringVarP(&author, "author", "a", "", "Author name")
	generateCmd.Flags().StringVarP(&author, "email", "e", "", "Author email")
	generateCmd.Flags().StringVarP(&out, "out", "o", "", "Output directory")
}
