package main

import (
	// Force import to test that everything builds correctly
	_ "github.com/pydio/cells-sdk-go/api"
	"github.com/pydio/cells-sdk-go/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
