package idm

import (
	"testing"

	"github.com/pydio/cells-sdk-go/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserService(t *testing.T) {

	Convey("Test simple CRUD", t, func() {

		if !config.RunEnvAwareTests {
			return // We skip tests that fail if we are not connected to a running Pydio Cells instance
		}

		idmUser, err := CreateUser("/testers", "user1", "Password123$")
		So(err, ShouldBeNil)
		So(idmUser.Login, ShouldEqual, "user1")

	})
}
