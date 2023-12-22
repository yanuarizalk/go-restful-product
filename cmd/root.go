/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-restful-product",
	Short: "Test application",
	Long:  `1 CRUD endpoint & CLI interface involving GORM, mysql, fiber, migration & seeder, swaggo, docker, unit & integration testing`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		os.Exit(0)
	}
}

func init() {

}
