package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	// Sample configuration files
	configFileName = "config.sample.json"
)

func TestMain(m *testing.M) {

	if !RunEnvAwareTests {
		os.Exit(0)
	}
	crp, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("cannot set up environment", err.Error())
	}
	crp = filepath.Dir(crp)
	SetUpEnvironment(GetDefaultConfigFiles(crp))
	os.Exit(m.Run())
}

func TestJWT(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		Convey("Is config correctly set", func() {
			// fmt.Printf("## DEFAULT CONF \n%s - %s - %s - %s - %s - %s - %v\n", DefaultConfig.Protocol, DefaultConfig.Url, DefaultConfig.ClientKey, DefaultConfig.ClientSecret, DefaultConfig.User, DefaultConfig.Password, DefaultConfig.SkipVerify)
			So(DefaultConfig.Url, ShouldNotBeEmpty)
			So(DefaultConfig.ClientKey, ShouldNotBeEmpty)
			So(DefaultConfig.ClientSecret, ShouldNotBeEmpty)
			So(DefaultConfig.User, ShouldNotBeEmpty)
			So(DefaultConfig.Password, ShouldNotBeEmpty)
			So(DefaultConfig.SkipVerify, ShouldNotBeEmpty)
		})

		jwt, err := retrieveToken(DefaultConfig)
		So(err, ShouldBeNil)
		fmt.Println("Retrieved JWT: " + jwt)
		So(len(jwt), ShouldBeGreaterThan, 256)
	})
}
