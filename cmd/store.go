/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gen-crud/stores"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"strings"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := stores.AddStore(args)
		println(c)

		dir := "stores"

		for _, arg := range args {
			if strings.HasPrefix(arg, "dir_") {
				dir = strings.ReplaceAll(arg, "dir_", "")
			}
		}

		modelName := args[0]

		fileName := fmt.Sprintf("%v.ts", strcase.ToLowerCamel(utils.ToPlural(modelName)))

		utils.WriteFile(dir, fileName, c)

	},
}

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
