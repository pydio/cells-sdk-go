package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RelengCmd is the main entry point to meta utils to manage this go SDK.
var RelengCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Release engineering controller for the Cells Go SDK",
	Long: `
	RelengCmd is the main entry point to meta utils to manage this go SDK.
	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// flags := RootCmd.PersistentFlags()
}
