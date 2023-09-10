/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/myyrakle/gopring/internal/command"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please specify the project name")
			return
		}

		packageName := args[0]

		command.New(packageName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
