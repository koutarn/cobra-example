/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.01"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long:  "All software has version.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s", version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
