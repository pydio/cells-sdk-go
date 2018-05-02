package util

import (
	"fmt"
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

	Convey("Test Config retrieval", t, func() {
		conf := getServerConfig()
		fmt.Printf("value: %v\n", conf.Get("pydio.grpc.activity", "driver").String("Default used"))

		// 		So(conf.Get("cert", "proxy", "ssl").(string), ShouldEqual, false)
	})
}
