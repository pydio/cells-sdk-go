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

			for i, ws := range wss.Workspaces {
				fmt.Printf("#%d - %s, %s\n", i, ws.Label, ws.Slug)
			}
		})

		// Following tests are *not* run unless below parameters are defined and valid.
		// We yet leave the code here as a sample for your convenience.
		// datasource, slug, label, desc := "localtestds1", "test-wksp-2", "Test Workspace 2", "Another test description"
		datasource, slug, label, desc := "", "", "", ""
		if len(datasource) == 0 {
			return
		}

		Convey("Can create workspace", func() {

			wss, err := ListWorkspaces()
			So(err, ShouldBeNil)
			initialNb := wss.Total

			ws, err := CreateWorkspace(datasource, slug, label, desc)
			So(err, ShouldBeNil)
			So(ws.Label, ShouldEqual, label)
			// So(ws.Slug, ShouldEqual, slug)

			wss, err = ListWorkspaces()
			So(wss.Total, ShouldEqual, initialNb+1)

			// Convey("Cannot create a second workspace with same slug", func() {
			// 	// FIXME: calling put on a CreateWorkspace with same slug creates a new workspace with suffixed slug (typically test-wksp-1)
			// 	ws, err = CreateWorkspace(datasource, label, slug)
			// 	So(err, ShouldBeNil)
			// 	So(ws.Label, ShouldEqual, label)
			// 	So(ws.Slug, ShouldNotEqual, slug)
			// })

			Convey("Can delete newly created workspace", func() {
				ok, err := DeleteWorkspace(slug)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			})

		})
	})
}
