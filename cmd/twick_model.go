package cmd

import (
	"path/filepath"

	"github.com/pydio/cells-sdk-go/utils"
	"github.com/spf13/cobra"
)

var (
	pathToModels string
)

var TwickCmd = &cobra.Command{
	Use:   "twick-model",
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

	TwickCmd.Flags().StringVar(&pathToModels, "path", "../models", "Path to the models package")

	RootCmd.AddCommand(TwickCmd)
}
