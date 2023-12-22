/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.yanuarizal.net/go-restful-product/database"
)

func showMigratesArg() {
	args := []string{}
	for arg := range database.MigrationList {
		args = append(args, arg)
	}

	fmt.Println("List of args: ", strings.Join(args, ", "))
}

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Execute a migration",
	Run: func(cmd *cobra.Command, args []string) {
		var exitCode int

		defer os.Exit(exitCode)

		if len(args) <= 0 {
			showMigratesArg()
			return
		}

		doMigrate := strings.TrimSpace(strings.ToLower(args[0]))

		if exec, exist := database.MigrationList[doMigrate]; exist {
			fmt.Println("migrating...")
			if err := exec(database.App, nil); err != nil {
				fmt.Println("migrating failed: ", err)
				exitCode = 1
				return
			}
			fmt.Println("successfully migrated")
		} else {
			showMigratesArg()
			exitCode = 1
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
