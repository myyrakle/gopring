/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/myyrakle/gopring/internal/generator"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build the project",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Generate()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
