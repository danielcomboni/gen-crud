/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gen-crud/models"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"strings"
)

// uimodelCmd represents the uimodel command
var uimodelCmd = &cobra.Command{
	Use:   "uimodel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := new(models.GenerateModelUI).SetModelArgs(args)
		println(c.GetModel())
		println(c.AddInstantiation())

		dir := "models"

		for _, arg := range args {
			if strings.HasPrefix(arg, "dir_") {
				dir = strings.ReplaceAll(arg, "dir_", "")
			}
		}

		d := fmt.Sprintf("%v\n%v", c.GetModel(), c.AddInstantiation())
		fileName := fmt.Sprintf("%v.ts", strcase.ToLowerCamel(c.GetModelName()))
		utils.WriteFile(dir, fileName, d)
	},
}

func init() {
	rootCmd.AddCommand(uimodelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uimodelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uimodelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
