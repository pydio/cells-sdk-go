package s3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/pydio/cells-sdk-go/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {

	if !config.RunEnvAwareTests {
		return
	}

	// Enhance this
	if data, e := ioutil.ReadFile("../../config.json"); e == nil {
		var c config.SdkConfig
		if e := json.Unmarshal(data, &c); e != nil {
			log.Fatal("Cannot decode config content", e)
		}
		config.DefaultConfig = &c
		fmt.Println("Connecting to " + config.DefaultConfig.Url)
	} else {
		log.Fatal("Cannot read file ", e)
	}

	if data, e := ioutil.ReadFile("../../config-s3.json"); e == nil {
		var c config.S3Config
		if e := json.Unmarshal(data, &c); e != nil {
			log.Fatal("Cannot decode s3 config content", e)
		}
		c.IsDebug = true
		config.DefaultS3Config = &c
		config.InitDefaultSession(c)
	} else {
		log.Fatal("Cannot read s3 config file ", e)
	}

	os.Exit(m.Run())
}

func TestS3Service(t *testing.T) {

	Convey("Test simple Get", t, func() {
		err := GetFile("", "", false)
		So(err, ShouldBeNil)
	})
}
