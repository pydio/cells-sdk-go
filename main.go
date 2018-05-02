package main

import (
	"fmt"

	// Force import to test that everything builds correctly
	_ "github.com/pydio/cells-sdk-go/api"

	"github.com/pydio/cells-sdk-go/example/idm"
)

func main() {
	fmt.Println("SDK loaded...")

	idm.CreateUser("/dev", "testDevUser3", "withAPassword")
	idm.ListUsers()

}
