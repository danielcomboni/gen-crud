/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gen-crud/httpclient"
	"gen-crud/utils"
	"github.com/spf13/cobra"
)

// httpclientCmd represents the httpclient command
var httpclientCmd = &cobra.Command{
	Use:   "httpclient",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _, content := httpclient.AddHttpClient(args)
		utils.WriteFile(dir, "handler.ts", content)
	},
}

func init() {
	rootCmd.AddCommand(httpclientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpclientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpclientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
