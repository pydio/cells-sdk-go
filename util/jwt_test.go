package util

import (
	"testing"

	"github.com/pydio/cells-sdk-go/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJWT(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		if !config.RunEnvAwareTests {
			return // We skip tests that fail if we are not connected to a running Pydio Cells instance
		}

		jwt, err := retrieveToken(adminUser, adminPwd)
		So(err, ShouldBeNil)
		So(len(jwt), ShouldBeGreaterThan, 400)
		// tokens := strings.Split(jwt, "-")
		// So(len(tokens), ShouldEqual, 3)
	})

}
