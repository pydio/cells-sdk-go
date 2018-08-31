package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/config"
)

var (
	configFile string

	protocol   string
	host       string
	id         string
	secret     string
	user       string
	pwd        string
	skipVerify bool
)

// RootCmd is the parent of all REST API Client commands.
// It encapsulates the mechanism to configure the communication with a running Cells server.
var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Simple CLI Client for Pydio Cells REST API",
	Long: `
# Simple CLI Client for Pydio Cells REST API.
	
You should either provide config parameters using the flags or a path to a valid json configuration file.
	
### Syntax

$ ` + os.Args[0] + ` [command] [flags] args...

### Examples

Use a json configuration file that lays in the same folder (see config/config.sample.json file for an example)  
$ ` + os.Args[0] + ` [command] --c config.json

Use parameters as flag 
$ ` + os.Args[0] + ` [command] -protocol http --u localhost:8080 --i pydio-front --s pydio-secret --l admin --p pass123 


	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Parse from parameters
		if host != "" && id != "" && secret != "" && user != "" && pwd != "" {
			config.DefaultConfig = &config.SdkConfig{
				Protocol:     protocol,
				Url:          host,
				ClientKey:    id,
				ClientSecret: secret,
				User:         user,
				Password:     pwd,
				SkipVerify:   skipVerify,
			}
			return
		}

		// Parse from config
		if configFile == "" {
			cmd.Help()
			cmd.Println("\n--\nError: please provide a path to the configuration file or provide config parameters (See above).")
			os.Exit(1)
		}

		data, e := ioutil.ReadFile(configFile)
		if e != nil {
			log.Fatal("cannot read config file:", e)
		}

		var c config.SdkConfig
		if e = json.Unmarshal(data, &c); e != nil {
			log.Fatal("Cannot decode config content for file at", configFile, "- route cause:", e)
		}
		config.DefaultConfig = &c

		// if data, e := ioutil.ReadFile(configFile); e == nil {
		// 	var c config.SdkConfig
		// 	if e := json.Unmarshal(data, &c); e != nil {
		// 		log.Fatal("Cannot decode config content", e)
		// 	}
		// 	config.DefaultConfig = &c
		// 	// fmt.Println("Connecting to " + config.DefaultConfig.Url)
		// 	// fmt.Println("")
		// } else {
		// 	log.Fatal("Cannot read file at", configFile, ", root cause:", e)
		// }

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "", "Relative path to the configuration file")

	flags.StringVar(&protocol, "protocol", "http", "Http scheme to server")
	flags.StringVarP(&host, "url", "u", "", "Http URL to server")
	flags.StringVarP(&id, "id", "i", "", "OIDC Client ID")
	flags.StringVarP(&secret, "secret", "s", "", "OIDC Client Secret")
	flags.StringVarP(&user, "login", "l", "", "User login")
	flags.StringVarP(&pwd, "password", "p", "", "User password")
	flags.BoolVar(&skipVerify, "skipVerify", false, "Skip SSL certificate verification (not recommended)")

}
