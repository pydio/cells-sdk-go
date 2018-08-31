package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	rootPath string
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestJWT(t *testing.T) {

	if RunEnvAwareTests {
		Convey("Given a clean environment than can access a running Cells instance", t, func() {

			Convey("Set up default env aware configuration", func() {
				crp, err := filepath.Abs(".")
				So(err, ShouldBeNil)

				rootPath = filepath.Dir(crp)
				SetUpEnvironment(GetDefaultConfigFiles(rootPath))

				Convey("Configuration must have been correctly set", func() {

					// fmt.Printf("## DEFAULT CONF \n%s - %s - %s - %s - %s - %s - %v\n", DefaultConfig.Protocol, DefaultConfig.Url, DefaultConfig.ClientKey, DefaultConfig.ClientSecret, DefaultConfig.User, DefaultConfig.Password, DefaultConfig.SkipVerify)
					So(DefaultConfig.Url, ShouldNotBeEmpty)
					So(DefaultConfig.ClientKey, ShouldNotBeEmpty)
					So(DefaultConfig.ClientSecret, ShouldNotBeEmpty)
					So(DefaultConfig.User, ShouldNotBeEmpty)
					So(DefaultConfig.Password, ShouldNotBeEmpty)
					So(DefaultConfig.SkipVerify, ShouldNotBeEmpty)
					sdkc, s3c := GetDefaultConfigFiles(rootPath)
					fmt.Println("Config corretly set from: ", sdkc, s3c)
					fmt.Println("Config Client Secret: ", DefaultConfig.ClientSecret)

				})

				Convey("JWT token must be retrievable and longer than 256 chars.", func() {
					jwt, err := retrieveToken(DefaultConfig)
					So(err, ShouldBeNil)
					So(len(jwt), ShouldBeGreaterThan, 256)
				})
			})
		})

	}

}
