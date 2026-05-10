/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"language-learning/internal/bootstrap"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "language-learning",
	Short: "language-learning backend",
	Long:  `basic run`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Bootstrap()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
