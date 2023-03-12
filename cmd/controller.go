// Package cmd /*

/*
Copyright © 2023 Daniel Comboni comboni93@gmail.com
*/

package cmd

import (
	"fmt"
	"gen-crud/controllers"
	"gen-crud/utils"
	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "generates controller for a particular model",
	Long:  `generates a controller for a particular model`,
	Run: func(cmd *cobra.Command, args []string) {
		c := new(controllers.GenerateController).SetControllerArgs(args)
		println(c.AddCRUD())
		utils.WriteFile("controllers", fmt.Sprintf("%v.go", c.GetControllerName()), c.AddCRUD())
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// controllerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
