package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJWT(t *testing.T) {

	Convey("Test JWT retrieval", t, func() {

		if !RunEnvAwareTests {
			return // We skip tests that fail if we are not connected to a running Pydio Cells instance
		}

		_, err := retrieveToken(&SdkConfig{})
		So(err, ShouldNotBeNil)
		// So(len(jwt), ShouldBeGreaterThan, 400)
		// tokens := strings.Split(jwt, "-")
		// So(len(tokens), ShouldEqual, 3)
	})

}
