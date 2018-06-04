package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {

	if !RunEnvAwareTests {
		os.Exit(0)
	}

	// Enhance this
	if data, e := ioutil.ReadFile("../config.json"); e == nil {
		var c SdkConfig
		if e := json.Unmarshal(data, &c); e != nil {
			log.Fatal("Cannot decode config content", e)
		}
		DefaultConfig = &c
		fmt.Println("Connecting to " + DefaultConfig.Url)
	} else {
		log.Fatal("Cannot read file ", e)
	}

	os.Exit(m.Run())
}

func TestJWT(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		Convey("Is config correctly set", func() {
			fmt.Printf("## DEFAULT CONF \n%s - %s - %s - %s - %s - %s - %v\n", DefaultConfig.Protocol, DefaultConfig.Url, DefaultConfig.ClientKey, DefaultConfig.ClientSecret, DefaultConfig.User, DefaultConfig.Password, DefaultConfig.SkipVerify)
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
