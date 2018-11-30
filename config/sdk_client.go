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
	apiResourcePath  = "/a"
	oidcResourcePath = "/auth/dex"

	grantType = "password"
	scope     = "email profile pydio"

	store = NewTokenStore()
)

// GetAnonymousApiClient prepares a client WITHOUT sending a first request to the OIDC service
func GetAnonymousApiClient(sdkConfig *SdkConfig) (context.Context, *apiclient.PydioCellsRest, error) {
	transport := httptransport.New(sdkConfig.Url, apiResourcePath, []string{sdkConfig.Protocol})
	client := apiclient.New(transport, strfmt.Default)
	return context.Background(), client, nil
}

// GetPreparedApiClient connects to the Pydio Cells server defined by this config, by sending an authentication
// request to the OIDC service to get a valid JWT (or taking the JWT from cache).
// Also returns a context to be used in subsequent requests.
func GetPreparedApiClient(sdkConfig *SdkConfig) (context.Context, *apiclient.PydioCellsRest, error) {

	transport := httptransport.New(sdkConfig.Url, apiResourcePath, []string{sdkConfig.Protocol})
	jwt, err := retrieveToken(sdkConfig)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"cannot retrieve token with config:\n%s - %s - %s - %s - %s - %s - %v\nerror cause: %s",
			sdkConfig.Protocol, sdkConfig.Url, sdkConfig.ClientKey, sdkConfig.ClientSecret,
			sdkConfig.User, sdkConfig.Password, sdkConfig.SkipVerify, err.Error())
	}
	bearerTokenAuth := httptransport.BearerToken(jwt)
	transport.DefaultAuthentication = bearerTokenAuth

	client := apiclient.New(transport, strfmt.Default)

	return context.Background(), client, nil
}

// GetHttpClient adds an option to rather use an http client that ignore SSL certificate issues.
func GetHttpClient(sdkConfig *SdkConfig) *http.Client {

	if sdkConfig.SkipVerify {
		log.Println("[WARNING] Using SkipVerify for ignoring SSL certificate issues!")
		return &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}}
	}
	return http.DefaultClient
}
