package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	regexpSampleFilePath = "./sample.txt"
)

func TestFindString(t *testing.T) {

	Convey("Sample file should be found", t, func() {
		_, e := ioutil.ReadFile(regexpSampleFilePath)
		So(e, ShouldBeNil)
	})

	Convey("Pattern must compile", t, func() {
		re := regexp.MustCompile(replacePattern)

		r, _ := ioutil.ReadFile(regexpSampleFilePath)

		matches := re.FindStringSubmatch(string(r))
		fmt.Printf("Done, found %d matches\n", len(matches))
		fmt.Printf("%q", matches)
	})

	Convey("Samples", t, func() {
		// Some example from go doc
		// re := regexp.MustCompile("a(x*)b")
		// fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
		// fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
		// fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
		// fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

		re := regexp.MustCompile(replacePattern)
		r, _ := ioutil.ReadFile(regexpSampleFilePath)
		initialContent := string(r)
		modifiedContent := re.ReplaceAllString(initialContent, "int64 `json:\"$1,string,omitempty\"`")
		So(initialContent, ShouldNotEqual, modifiedContent)
	})
}
