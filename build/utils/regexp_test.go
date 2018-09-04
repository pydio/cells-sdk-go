package utils

import (
	"io/ioutil"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	regexpSampleFilePath = "./sample.txt"
)

func TestFindString(t *testing.T) {

	Convey("Given a simple sample file", t, func() {
		r, e := ioutil.ReadFile(regexpSampleFilePath)
		So(e, ShouldBeNil)

		Convey("Pattern must compile", func() {
			re := regexp.MustCompile(replacePattern)

			Convey("2 matches must be found", func() {
				matches := re.FindStringSubmatch(string(r))
				So(len(matches), ShouldEqual, 2)

				// fmt.Printf("Done, found %d matches\n", )
				// fmt.Printf("%q", matches)
			})

			Convey("Source string should be modified", func() {
				initialContent := string(r)
				modifiedContent := re.ReplaceAllString(initialContent, "int64 `json:\"$1,string,omitempty\"`")
				So(initialContent, ShouldNotEqual, modifiedContent)
			})
		})
	})

}
