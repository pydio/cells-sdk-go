package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	cells_sdk "github.com/pydio/cells-sdk-go"
)

const (
	sampleFilePath   = "config.sample.json"
	sampleS3FilePath = "config-s3.sample.json"
)

func TestSetUpProcess(t *testing.T) {

	Convey("Given a clean environment:", t, func() {

		Convey("Sample files should be found", func() {
			_, e := ioutil.ReadFile(sampleFilePath)
			So(e, ShouldBeNil)

			_, e = ioutil.ReadFile(sampleS3FilePath)
			So(e, ShouldBeNil)
		})

		// We run subsequent tests only in basic builds:
		// It breaks other tests when in EnvAware mode because initialisation of the env is kept by convey
		Convey("With an unconfigured environment", func() {

			Convey("Sample config files should be correctly formatted", func() {
				s, _ := ioutil.ReadFile(sampleFilePath)
				var c cells_sdk.SdkConfig
				e := json.Unmarshal(s, &c)
				So(e, ShouldBeNil)
				So(c.User, ShouldEqual, "admin")

				s3, _ := ioutil.ReadFile(sampleS3FilePath)
				var c2 cells_sdk.S3Config
				e = json.Unmarshal(s3, &c2)
				So(e, ShouldBeNil)
				So(c2.ApiSecret, ShouldEqual, "gatewaysecret")
			})

			Convey("Sample SDK config should be correcly parsed and loaded", func() {
				s, _ := ioutil.ReadFile(sampleFilePath)
				var c cells_sdk.SdkConfig
				json.Unmarshal(s, &c)
				So(c.User, ShouldEqual, "admin")

				DefaultConfig = &c
				So(DefaultConfig.User, ShouldEqual, "admin")
			})

			Convey("Sample S3 Config should be parsed and loaded", func() {
				s3, _ := ioutil.ReadFile(sampleS3FilePath)
				var c cells_sdk.S3Config
				e := json.Unmarshal(s3, &c)
				So(e, ShouldBeNil)

				DefaultS3Config = &c
				So(DefaultS3Config.ApiSecret, ShouldEqual, "gatewaysecret")
			})
		})

	})

	Convey("Given a clean environment", t, func() {

		Convey("Cells SDK must be configured", func() {
			var c cells_sdk.SdkConfig
			// check presence of Env variable
			url := os.Getenv(KeyURL)
			clientKey := os.Getenv(KeyClientKey)
			clientSecret := os.Getenv(KeyClientSecret)
			user := os.Getenv(KeyUser)
			password := os.Getenv(KeyPassword)
			skipVerifyStr := os.Getenv(KeySkipVerify)
			if skipVerifyStr == "" {
				skipVerifyStr = "false"
			}
			skipVerify, err := strconv.ParseBool(skipVerifyStr)
			So(err, ShouldBeNil)
			if len(url) > 0 && len(clientKey) > 0 && len(clientSecret) > 0 && len(user) > 0 && len(password) > 0 {
				c.Url = url
				c.ClientKey = clientKey
				c.ClientSecret = clientSecret
				c.User = user
				c.Password = password
				c.SkipVerify = skipVerify
			} else { // fallback to config.json file
				currPath, err := filepath.Abs(".")
				So(err, ShouldBeNil)
				configFilePath := filepath.Join(currPath, sampleFilePath)
				s, err := ioutil.ReadFile(configFilePath)
				e := json.Unmarshal(s, &c)
				So(e, ShouldBeNil)
			}

			DefaultConfig = &c
			So(DefaultConfig.User, ShouldEqual, "admin")
		})
	})

}
