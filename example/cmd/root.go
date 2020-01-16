// Package cmd implements some basic examples of what can be achieved when combining
// the use of the Go SDK for Cells with the powerful Cobra framework to implement CLI
// client applications for Cells.
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	cells_sdk "github.com/pydio/cells-sdk-go"
)

var (
	configFile string

	host       string
	id         string
	secret     string
	user       string
	pwd        string
	skipVerify bool
)

// ExampleCmd is the parent of all example commands defined in this package.
// It takes care of the pre-configuration of the defaut connection to the SDK
// in its PersistentPreRun phase.
var ExampleCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Sample commands to show how to use the Go SDK for Pydio Cells",
	Long: `
# Sample commands to show how to use the Go SDK for Pydio Cells

Pydio Cells comes with a powerful REST API that exposes various endpoints and enable management of a running Cells instance.
As a convenience, the Pydio team also provide a ready to use SDK for the Go language that encapsulates the boiling code to wire things 
and provides a few chosen utilitary methods to ease implemantation when using the SDK in various Go programs.

The children commands defined here show some basic examples of what can be achieved when combining the use of this SDK with 
the powerful Cobra framework to easily implement small CLI client applications.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if configFile == "" {
			// Parse from parameters

			// insure all necessary parameters are defined
			var msg string
			if host == "" {
				msg += " - Host URL\n"
			}
			if user == "" {
				msg += " - the login of an existing user\n"
			}
			if pwd == "" {
				msg += " - the password of an existing user\n"
			}

			if len(msg) > 0 {
				cmd.Println("Could not set up default configuration, missing arguments:\n", msg,
					"\n\nYou might also directly provide the relative path to a config.json file. See in-line help for more details.")
				os.Exit(1)
			}

			DefaultConfig = &cells_sdk.SdkConfig{
				Url:        host,
				User:       user,
				Password:   pwd,
				SkipVerify: skipVerify,
			}

			return
		}

		data, e := ioutil.ReadFile(configFile)
		if e != nil {
			log.Fatal("cannot read config file:", e)
		}

		var c cells_sdk.SdkConfig
		if e = json.Unmarshal(data, &c); e != nil {
			log.Fatal("Cannot decode config content for file at", configFile, "- route cause:", e)
		} else {
			DefaultConfig = &c
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	flags := ExampleCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "", "Path to the configuration file")

	flags.StringVarP(&host, "url", "u", "", "HTTP URL to server")
	flags.StringVarP(&id, "api-key", "k", "", "OIDC Client ID")
	flags.StringVarP(&secret, "api-secret", "s", "", "OIDC Client Secret")
	flags.StringVarP(&user, "login", "l", "", "User login")
	flags.StringVarP(&pwd, "password", "p", "", "User password")
	flags.BoolVar(&skipVerify, "skipVerify", false, "Skip SSL certificate verification (not recommended)")

}
