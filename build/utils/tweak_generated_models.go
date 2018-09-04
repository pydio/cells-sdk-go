package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var (
	replacePattern = "int64 `json:\"([a-zA-Z_][a-zA-Z_0-9]*),omitempty\"`"
	replacer       = regexp.MustCompile(replacePattern)
)

// AdaptGeneratedModels modify the go files generated with go-swagger to enable communication with Pydio Cells' REST API
func AdaptGeneratedModels(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil
	}

	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		fmt.Println("Modifying: " + path)

		newContents := replacer.ReplaceAllString(string(read), "int64 `json:\"$1,string,omitempty\"`")

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			return err
		}
	}
	return nil
}
