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

		initialNb := dss.Total

		So(dss.Total, ShouldNotEqual, 0)

		for i, ds := range dss.DataSources {
			fmt.Printf("#%d - %s\n", i, ds.Name)
		}

		nds, err := AddLocalDatasource("pydiodstest1", "192.168.0.140", 9009, "/home/bsinou/tmp/pydio/dss/")
		So(err, ShouldBeNil)
		fmt.Printf("Created data source: %s\n", nds.Name)

		dss, err = ListDatasources()
		So(dss.Total, ShouldEqual, initialNb+1)
		for i, ds := range dss.DataSources {
			fmt.Printf("#%d - %s\n", i, ds.Name)
		}

		ok, err := DeleteLocalDatasource("pydiodstest1")
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)

		dss, err = ListDatasources()
		So(dss.Total, ShouldEqual, initialNb)

	})
}
