package s3

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	config.SetUpEnvironment(config.GetDefaultConfigFiles(crp))
	os.Exit(m.Run())
}

func TestS3Service(t *testing.T) {

	Convey("CRUD via s3 works", t, func() {

		wkspSlug := "personal-files"
		testFileName := "test.txt"
		testStr := "Simple test content"
		testStr2 := testStr + " updated"

		// Create
		_, err := PutFile(wkspSlug, testFileName, strings.NewReader(testStr))
		So(err, ShouldBeNil)
		time.Sleep(1000 * time.Millisecond)

		// Read
		output, err := GetFile(wkspSlug, testFileName)
		So(err, ShouldBeNil)
		buf := new(bytes.Buffer)
		i, err := buf.ReadFrom(output.Body)
		retrievedStr := buf.String()
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 19)
		So(retrievedStr, ShouldEqual, testStr)

		// Update
		// fmt.Println("About to update with content: " + testStr2)
		_, err = PutFile(wkspSlug, testFileName, strings.NewReader(testStr2))
		So(err, ShouldBeNil)

		// for j := 0; j < 10; j++ {
		time.Sleep(1000 * time.Millisecond)
		output, err = GetFile(wkspSlug, testFileName)
		So(err, ShouldBeNil)
		buf = new(bytes.Buffer)
		i, err = buf.ReadFrom(output.Body)
		retrievedStr2 := buf.String()
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 27)

		// fmt.Printf("retrieved length: %d \n", i)

		// fmt.Printf("#%d: retrieved: [%s] \n", j, retrievedStr2)
		So(retrievedStr2, ShouldEqual, testStr2)
		// }

		// Delete
		_, err = DeleteFile(wkspSlug, testFileName)
		So(err, ShouldBeNil)

		// Try to retrieve deleted object
		obj, err := GetFile(wkspSlug, testFileName)
		So(err, ShouldNotBeNil)
		So(obj, ShouldBeNil)

	})

	// Convey("Test simple Get", t, func() {
	// 	err := GetFile("basic", "test/four.jpeg", false)
	// 	So(err, ShouldBeNil)
	// })

}
