package tree

import (
	"fmt"
	"testing"

	"github.com/pydio/cells-sdk-go/config"
	. "github.com/smartystreets/goconvey/convey"
)

// See datasources_test.go
// func TestMain(m *testing.M) {}

func TestWorkspaceService(t *testing.T) {

	Convey("Given a correctly setup environment", t, func() {
		So(config.RunEnvAwareTests, ShouldBeTrue) // Or should never reach this point otherwise

		So(config.DefaultConfig, ShouldNotBeNil)
		So(len(config.DefaultConfig.Url), ShouldBeGreaterThan, 0)

		Convey("Can list workspaces", func() {

			wss, err := ListWorkspaces()
			So(err, ShouldBeNil)
			fmt.Printf(" (found %d)", wss.Total)

			Convey("We have at least one workspace", func() {
				So(wss.Total, ShouldBeGreaterThanOrEqualTo, 1)
			})
		})

		// Following tests are not run unless below parameters are defined and valid. We yet leave the code here as a sample for your convenience.
		// datasource, label, slug  := "localtestds1", "Test Workspace", "test-wksp"
		datasource, label, slug := "", "", ""

		if len(datasource) == 0 {
			return
		}

		ws, err := CreateWorkspace(datasource, label, slug)
		So(err, ShouldBeNil)
		fmt.Printf("Created workspace: %s\n", ws.Label)

		// dss, err = ListDatasources()
		// So(dss.Total, ShouldEqual, initialNb+1)
		// for i, ds := range dss.DataSources {
		// 	fmt.Printf("#%d - %s\n", i, ds.Name)
		// }

		// ok, err := DeleteLocalDatasource("pydiodstest1")
		// So(err, ShouldBeNil)
		// So(ok, ShouldBeTrue)

		// dss, err = ListDatasources()
		// So(dss.Total, ShouldEqual, initialNb)
	})
}
