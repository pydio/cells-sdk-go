package main

import (
	"log"

	"github.com/pydio/cells-sdk-go/v5/apiv1/example/cmd"
)

func main() {
	err := cmd.ExampleCmd.Execute()
	if err != nil {
		log.Fatal("could not launch example command:", err)
	}
}
