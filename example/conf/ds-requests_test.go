package conf

import (
	"fmt"
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
	err = config.SetUpEnvironment(config.GetDefaultConfigFiles(crp))
	if err != nil {
		fmt.Println("cannot set up environment", err.Error())
	}
	os.Exit(m.Run())
}

func TestSetup(t *testing.T) {

	Convey("Test Default Config", t, func() {
		So(config.RunEnvAwareTests, ShouldBeTrue) // Or should never reach this point otherwise
		// So(config.ConfigFileName, ShouldEqual, "../config.json")
	})
}

func TestDSService(t *testing.T) {

	Convey("Test simple datasource list", t, func() {
		So(config.RunEnvAwareTests, ShouldBeTrue) // Or should never reach this point otherwise

		So(config.DefaultConfig, ShouldNotBeNil)
		So(len(config.DefaultConfig.Url), ShouldBeGreaterThan, 0)

		dss, err := ListDatasources()
		So(err, ShouldBeNil)

		So(dss.Total, ShouldNotEqual, 0)
	})
}
