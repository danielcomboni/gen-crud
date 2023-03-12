// Package cmd /*

/*
Copyright © 2023 Daniel Comboni comboni93@gmail.com
*/

package cmd

import (
	"errors"
	"fmt"
	"gen-crud/models"
	"gen-crud/utils"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
	//"strings"
)

// modelCmd represents the hello command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "creates a model with fields",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(fmt.Sprintf("start model generation: %v", args))

		m := new(models.GenerateModel).SetModelArgs(args).SetStruct()
		println(m.GetStruct())

		var dirs []string
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			dirs = append(dirs, f.Name())
		}

		if _, err := os.Stat(m.GetPackageName()); errors.Is(err, os.ErrNotExist) {
			log.Println("creating directory...")
			err := os.Mkdir(m.GetPackageName(), os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		utils.WriteFile(m.GetPackageName(), fmt.Sprintf("%v.go", strings.ToLower(strcase.ToSnake(m.GetModelName()))), m.GetStruct())
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
	modelCmd.LocalFlags().String("req", "", "set required to validate a field")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
