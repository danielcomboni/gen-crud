// Package cmd /*

/*
Copyright © 2023 Daniel Comboni comboni93@gmail.com
*/

package cmd

import (
	"fmt"
	"gen-crud/routes"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := new(routes.GenerateRoute).SetArgs(args)
		println(r.GenerateRouteString())
		f := fmt.Sprintf("%v_routes", strcase.ToSnake(r.ModelName))
		//utils.WriteFile("routes", "")
		println(f)
		utils.WriteFile("routes", fmt.Sprintf("%v_routes.go", strcase.ToSnake(r.ModelName)), r.GenerateRouteString())
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
