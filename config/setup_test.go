package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	sampleFilePath   = "config.sample.json"
	sampleS3FilePath = "config-s3.sample.json"
)

func TestSetUpProcess(t *testing.T) {

	Convey("Sample files should be found", t, func() {
		_, e := ioutil.ReadFile(sampleFilePath)
		So(e, ShouldBeNil)

		_, e = ioutil.ReadFile(sampleS3FilePath)
		So(e, ShouldBeNil)
	})

	Convey("Environment Config must be loaded", t, func() {
		s, _ := ioutil.ReadFile(sampleFilePath)
		var c SdkConfig
		e := json.Unmarshal(s, &c)
		So(e, ShouldBeNil)
		So(c.User, ShouldEqual, "admin")

		s3, _ := ioutil.ReadFile(sampleS3FilePath)
		var c2 S3Config
		e = json.Unmarshal(s3, &c2)
		So(e, ShouldBeNil)
		So(c2.ApiSecret, ShouldEqual, "gatewaysecret")
	})

	Convey("SDK should be configured", t, func() {
		s, _ := ioutil.ReadFile(sampleFilePath)
		var c SdkConfig
		json.Unmarshal(s, &c)
		So(c.User, ShouldEqual, "admin")

		DefaultConfig = &c
		So(DefaultConfig.User, ShouldEqual, "admin")
	})

	Convey("Default S3Config should be loaded", t, func() {
		s3, _ := ioutil.ReadFile(sampleS3FilePath)
		var c S3Config
		e := json.Unmarshal(s3, &c)
		So(e, ShouldBeNil)

		DefaultS3Config = &c
		So(DefaultS3Config.ApiSecret, ShouldEqual, "gatewaysecret")
	})
}

func TestEnvConfiguration(t *testing.T) {

	Convey("Environment must be configured", t, func() {

		Convey("Pydio Cells SDK must be configured", func() {
			var c SdkConfig
			// check presence of Env variable
			protocol := os.Getenv(KeyProtocol)
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
			if len(protocol) > 0 && len(url) > 0 && len(clientKey) > 0 && len(clientSecret) > 0 && len(user) > 0 && len(password) > 0 {
				c.Protocol = protocol
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

			// Convey("Target Server must be reachable", func() {
			// 	_, _, err := GetPreparedApiClient(DefaultConfig)
			// 	So(err, ShouldBeNil)
			// })
		})
	})

	// Convey("Setup must be complete", t, func() {
	// 	err := SetUpEnvironment()
	// 	So(err, ShouldBeNil)
	// 	So(DefaultConfig.User, ShouldEqual, "admin")
	// 	So(DefaultS3Config.Bucket, ShouldEqual, "io")
	// })
}

// func TestSetUp(t *testing.T) {

// 	Convey("Target Server must be reachable", t, func() {
// 		SetUpEnvironment()
// 		_, _, err := GetPreparedApiClient(DefaultConfig)
// 		So(err, ShouldBeNil)
// 	})

// }
