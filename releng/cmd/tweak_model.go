package cmd

import (
	"path/filepath"

	"github.com/pydio/cells-sdk-go/releng/utils"
	"github.com/spf13/cobra"
)

var (
	pathToModels string
)

// TweakCmd launch automated manipulation of the generated model to workaround some known issues.
var TweakCmd = &cobra.Command{
	Use:   "tweak-model",
	Short: "Modify go-swagger generated models to enable communication with the protobuf generated RESTful API",
	Run: func(cmd *cobra.Command, args []string) {

		err := filepath.Walk("./models", utils.AdaptGeneratedModels)
		if err != nil {
			cmd.Print("could not modify models: " + err.Error())
			return
		}

		cmd.Print("Success: files have been modified")
	},
}

func init() {

	TweakCmd.Flags().StringVar(&pathToModels, "path", "../models", "Path to the models package")

	RelengCmd.AddCommand(TweakCmd)
}
