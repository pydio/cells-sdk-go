package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pydio/cells-sdk-go/client/meta_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

var (
	protocol   = "http"
	host       = "192.168.0.171"
	id         = "cells-front"
	user       = "elsa"
	pwd        = "hosk"
	skipVerify bool
	secret     = "fc3dLqcTBIu3IN7iYKpvnfIQ"
)

func main() {
	sdkConfig := &config.SdkConfig{
		Protocol:     protocol,
		Url:          host,
		ClientKey:    id,
		ClientSecret: secret,
		User:         user,
		Password:     pwd,
		SkipVerify:   skipVerify,
	}
	// will try to connect 10 times with the wrong password on user defined above

	for i := 1; i <= 10; i++ {

		sdkConfig.Password = "thisisawrongpassword"

		_, _, err := config.GetPreparedApiClient(sdkConfig)

		if err != nil {
			log.Printf("\n apiClient ERROR -> %v \n %v", i, err)
			time.Sleep(500 * time.Millisecond)
		}
	}
	//11TH try to connect with right credentials and list files should not be able, will throw "User elsa is not authorized to log in"

	fmt.Println("11TH TRY HERE W/ RIGHT CRED")

	sdkConfig.Password = pwd
	_, _, err := config.GetPreparedApiClient(sdkConfig)

	listFile(sdkConfig)

	if err != nil {
		log.Printf("\n apiClient ERROR 11th TRY -> %v \n %v ", 11, err)
	}

}

func listFile(sdkConfig *config.SdkConfig) {

	httpClient := config.GetHttpClient(sdkConfig)
	apiClient, ctx, err := config.GetPreparedApiClient(sdkConfig)
	if err != nil {
		log.Fatal("fatal listfiles ", err)
	}

	params := &meta_service.GetBulkMetaParams{
		Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
			"/*",
			"personal-files/*",
		}},
		Context:    ctx,
		HTTPClient: httpClient,
	}

	result, err := apiClient.MetaService.GetBulkMeta(params)
	if err != nil {
		fmt.Printf("could not list meta: %s\n", err.Error())
		log.Fatal(err)
	}

	if len(result.Payload.Nodes) > 0 {
		fmt.Printf("* %d meta\n", len(result.Payload.Nodes))
		for _, u := range result.Payload.Nodes {
			fmt.Println("  -", u.Path)
		}
	}
}
