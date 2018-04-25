package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	username = "admin"
	password = "a super admin password"
)

func TestEmptyDao(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		jwt, err := retrieveToken(username, password)
		So(err, ShouldBeNil)
		So(len(jwt), ShouldBeGreaterThan, 400)
		// tokens := strings.Split(jwt, "-")
		// So(len(tokens), ShouldEqual, 3)
	})
}
