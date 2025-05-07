package main

import (
	"github.com/pydio/cells-sdk-go/v4/apiv1/example/cmd"
	"log"
)

func main() {
	err := cmd.ExampleCmd.Execute()
	if err != nil {
		log.Fatal("could not launch example command:", err)
	}
}
