package tree

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

func TestDatasourceService(t *testing.T) {

	Convey("Given a setup environment", t, func() {
		So(config.RunEnvAwareTests, ShouldBeTrue) // we should never reach this point otherwise
		So(config.DefaultConfig, ShouldNotBeNil)
		So(len(config.DefaultConfig.Url), ShouldBeGreaterThan, 0)

		Convey("Can List datasources", func() {

			dss, err := ListDatasources()
			So(err, ShouldBeNil)
			So(dss, ShouldNotBeNil)
			So(dss.Total, ShouldNotBeNil)

			Convey("Returned collection contains at least one datasource", func() {
				fmt.Printf("(found %d dss) ", dss.Total) // add details for the verbose mode
				So(dss.Total, ShouldBeGreaterThanOrEqualTo, 1)
			})

		})

		// Following tests are not run unless below parameters are defined and valid. We yet leave the code here as a sample for your convenience.
		name, peerAddress, dsPort, rootFolder := "localtestds1", "192.168.0.140", 9009, "/tmp/pydio/dss/unittest/localtestds1"
		// name, peerAddress, dsPort, rootFolder := "", "", 9009, ""

		if len(rootFolder) == 0 {
			return
		}

		Convey("CRUD operations work", func() {

			dss, err := ListDatasources()
			So(err, ShouldBeNil)
			initialNb := dss.Total

			Convey("Can add a local datasource", func() {

				nds, err := AddLocalDatasource(name, peerAddress, int32(dsPort), rootFolder)
				So(err, ShouldBeNil)
				So(nds, ShouldNotBeNil)

				Convey("Datasource is correctly created", func() {
					So(nds.Name, ShouldEqual, name)
					dss, err = ListDatasources()
					So(dss.Total, ShouldEqual, initialNb+1)
				})

				Convey("Datasource can be deleted", func() {
					ok, err := DeleteLocalDatasource("pydiodstest1")
					So(err, ShouldBeNil)
					So(ok, ShouldBeTrue)

					Convey("Collection list has been updated", func() {
						dss, err = ListDatasources()
						So(dss.Total, ShouldEqual, initialNb-1)
					})
				})
			})
		})
	})
}
