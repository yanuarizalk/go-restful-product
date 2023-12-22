/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.yanuarizal.net/go-restful-product/database"
	"github.yanuarizal.net/go-restful-product/database/seeder"
)

func showSeedsArg() {
	args := []string{}
	for arg := range database.SeederList {
		args = append(args, arg)
	}

	fmt.Println("List of args: ", strings.Join(args, ", "))
}

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Fill yer database with fake data",
	Run: func(cmd *cobra.Command, args []string) {
		var exitCode int

		defer os.Exit(exitCode)

		if len(args) <= 0 {
			showSeedsArg()
			return
		}

		doMigrate := strings.TrimSpace(strings.ToLower(args[0]))

		if exec, exist := database.SeederList[doMigrate]; exist {
			fmt.Println("seeding...")

			opt := seeder.Option{
				Count: 1,
			}

			if len(args) >= 2 {
				opt.Count, _ = strconv.Atoi(args[1])
			}

			if err := exec(database.App, &opt); err != nil {
				fmt.Println("seeding failed: ", err)
				exitCode = 1
				return
			}
			fmt.Println("successfully seeded")
		} else {
			showSeedsArg()
			exitCode = 1
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
