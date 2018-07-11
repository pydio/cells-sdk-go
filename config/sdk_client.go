package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/pydio/cells-sdk-go/client"
)

var (
	ApiResourcePath  = "/a"
	OidcResourcePath = "/auth/dex"

	grantType = "password"
	scope     = "email profile pydio"

	store = NewTokenStore()
)

// GetPreparedApiClient connects to the Pydio Cells server defined by this config.
// Also returns a context to be used in subsequent requests.
func GetPreparedApiClient(sdkConfig *SdkConfig) (*apiclient.PydioCellsRest, context.Context, error) {

	transport := httptransport.New(sdkConfig.Url, ApiResourcePath, []string{sdkConfig.Protocol})
	jwt, err := retrieveToken(sdkConfig)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"cannot retrieve token with config:\n%s - %s - %s - %s - %s - %s - %v\nerror cause: %s",
			DefaultConfig.Protocol, DefaultConfig.Url, DefaultConfig.ClientKey, DefaultConfig.ClientSecret,
			DefaultConfig.User, DefaultConfig.Password, DefaultConfig.SkipVerify, err.Error())
	}
	bearerTokenAuth := httptransport.BearerToken(jwt)
	transport.DefaultAuthentication = bearerTokenAuth

	client := apiclient.New(transport, strfmt.Default)

	return client, context.Background(), nil
}

func GetHttpClient(sdkConfig *SdkConfig) *http.Client {

	if sdkConfig.SkipVerify {
		log.Println("[WARNING] Using SkipVerify for ignoring SSL certificate issues!!")
		return &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}}
	}
	return http.DefaultClient
}
