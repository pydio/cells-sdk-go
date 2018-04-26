package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmptyDao(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		jwt, err := retrieveToken(adminUser, adminPwd)
		So(err, ShouldBeNil)
		So(len(jwt), ShouldBeGreaterThan, 400)
		// tokens := strings.Split(jwt, "-")
		// So(len(tokens), ShouldEqual, 3)
	})
}
