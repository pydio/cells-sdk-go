package util

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {

	Convey("Test Config retrieval", t, func() {

		// This retrieves the path of the *executable*
		// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		// We rather need the path to the source file for unit tests
		_, filename, _, _ := runtime.Caller(0)
		currDir := filepath.Dir(filename)
		So(filepath.Base(currDir), ShouldEqual, "util")

		dirPath, err := filepath.Abs(currDir)
		So(err, ShouldBeNil)
		pathToJSON := filepath.Join(dirPath, "dummydata", "sampleConfig.json")

		_, err = os.Stat(pathToJSON)
		So(err, ShouldBeNil)

		pydioConfigFilePath = pathToJSON

		v1 := GetConfigValue("services", "consul", "data_dir").String("Path not found, using default...")
		So(v1, ShouldEqual, "/tmp/consul")

		v2 := GetConfigValue("defaults", "datasource").String("Path not found, using default...")
		So(v2, ShouldEqual, "Path not found, using default...")
	})
}
