package idm

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

	os.Exit(m.Run())
}

func TestUserService(t *testing.T) {

	Convey("Test simple CRUD", t, func() {

		if !config.RunEnvAwareTests {
			return // We skip tests that fail if we are not connected to a running Pydio Cells instance
		}

		idmUser, err := CreateUser("/testers", "user1", "Password123$", false)
		So(err, ShouldBeNil)
		So(idmUser.Login, ShouldEqual, "user1")
		So(idmUser.Password, ShouldBeEmpty)
	})
}
