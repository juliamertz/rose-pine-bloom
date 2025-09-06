package cmd

import (
	"fmt"

	"github.com/rose-pine/rose-pine-bloom/docs"
	"github.com/rose-pine/rose-pine-bloom/version"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Project initialization",
	Long:  `Creates files to start a new Rosé Pine theme project.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: the command is not clear at this point
		buildCmd := "bloom TODO"
		if err := docs.EnsureReadmeWithBuildCommand(buildCmd, version.GetCurrentVersion()); err != nil {
			fmt.Println("unable to update README:", err)
		}

		if err := docs.EnsureLicense(); err != nil {
			fmt.Println("unable to update LICENSE:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
