package cmd

import (
	"encoding/json"
	"fmt"
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

var ExampleCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Rest Client of Pydio Cells",
	Long:  `Pydio Cells client for managing API`,
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

		//checks if the parameters are set when using the command without a config.json file
		if protocol == "" {
			log.Fatal("Provide the protocol type")
		}
		if host == "" {
			log.Fatal("Provide the host")
		}
		if id == "" {
			log.Fatal("Provide the id")
		}
		if user == "" {
			log.Fatal("Provide the user")
		}
		if pwd == "" {
			log.Fatal("Provide the password")
		}
		if secret == "" {
			log.Fatal("Provide a secert key")
		}

		// Parse from config
		if configFile == "" {
			log.Fatal("Please provide a path to the configuration file")
		}
		if data, e := ioutil.ReadFile(configFile); e == nil {
			var c config.SdkConfig
			if e := json.Unmarshal(data, &c); e != nil {
				log.Fatal("Cannot decode config content", e)
			}
			config.DefaultConfig = &c
			fmt.Println("Connecting to " + config.DefaultConfig.Url)
			fmt.Println("")
		} else {
			log.Fatal("Cannot read file, root cause:", e)
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	flags := ExampleCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "config.json", "Path to the configuration file")

	flags.StringVar(&protocol, "protocol", "http", "Http scheme to server")
	flags.StringVarP(&host, "url", "u", "", "Http URL to server")
	flags.StringVarP(&id, "id", "i", "", "OIDC Client ID")
	flags.StringVarP(&secret, "secret", "s", "", "OIDC Client Secret")
	flags.StringVarP(&user, "login", "l", "", "User login")
	flags.StringVarP(&pwd, "password", "p", "", "User password")
	flags.BoolVar(&skipVerify, "skipVerify", false, "Skip SSL certificate verification (not recommended)")

}
