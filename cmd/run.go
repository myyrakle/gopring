/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/myyrakle/gopring/internal/command"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "build and run the project",
	Run: func(cmd *cobra.Command, args []string) {
		command.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
