package cmd

import (
	"fmt"
	"os"

	"github.com/rose-pine/rose-pine-bloom/builder"
	"github.com/rose-pine/rose-pine-bloom/config"
	"github.com/spf13/cobra"
)

var deriveCmd = &cobra.Command{
	Use:   "derive [flags] <input>",
	Short: "Builds the template file from a given theme file",
	Long:  `Builds the template file from a given theme file.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateFormat(format)
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := builder.DeriveTemplate(&config.BuildTemplateConfig{
			Input:  args[0],
			Output: outputDir,
			Prefix: prefix,
			Format: format,
			Plain:  plain,
			Commas: !noCommas,
			Spaces: !noSpaces,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deriving theme: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	deriveCmd.Flags().StringVarP(&outputDir, "output", "o", "dist", "Directory for generated files")
	deriveCmd.Flags().StringVarP(&prefix, "prefix", "p", "$", "Color variable prefix")
	deriveCmd.Flags().StringVarP(&format, "format", "f", "hex", formatFlagUsage())
	deriveCmd.Flags().BoolVar(&plain, "plain", false, "Remove decorators from color values")
	deriveCmd.Flags().BoolVar(&noCommas, "no-commas", false, "Remove commas from color values")
	deriveCmd.Flags().BoolVar(&noSpaces, "no-spaces", false, "Remove spaces from color values")
	rootCmd.AddCommand(deriveCmd)
}
