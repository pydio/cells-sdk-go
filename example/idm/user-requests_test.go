package idm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserService(t *testing.T) {

	Convey("Test simple crud", t, func() {

		idmUser, err := CreateUser("/testers", "user1", "Password123$")
		So(err, ShouldBeNil)
		So(idmUser.Login, ShouldEqual, "user1")

	})
}
