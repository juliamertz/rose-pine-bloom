package cmd

import (
	"os"

	"github.com/rose-pine/rose-pine-bloom/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bloom",
	Short: "The Rosé Pine theme generator",
	Long:  `The Rosé Pine theme generator.`,
}

func Execute() {
	rootCmd.Version = version.GetCurrentVersion()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

