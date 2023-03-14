/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gen-crud/ui"
	"gen-crud/utils"
	"github.com/spf13/cobra"
	"strings"
)

// uicreateCmd represents the uicreate command
var uicreateCmd = &cobra.Command{
	Use:   "uicreate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1 - fileName
		// 2 - name of model to create
		// 3 - f_input-text-id-name

		path := args[0]
		fileName := args[1]

		if !strings.Contains(fileName, ".svelte") {
			fileName = fmt.Sprintf("%v.svelte", fileName)
		}

		c := ui.GenerateCreateSvelteComponent(args)
		utils.WriteFile(fmt.Sprintf(path), fileName, c)
	},
}

func init() {
	rootCmd.AddCommand(uicreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uicreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uicreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
