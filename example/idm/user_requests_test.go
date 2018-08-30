package idm

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/pydio/cells-sdk-go/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {

	if !config.RunEnvAwareTests {
		return
	}
	crp, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("cannot set up environment", err.Error())
	}
	crp = filepath.Dir(filepath.Dir(crp))
	config.SetUpEnvironment(config.GetDefaultConfigFiles(crp))
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
