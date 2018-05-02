package util

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	sdkapi "github.com/pydio/cells-sdk-go/api"
)

var (
	hostURL         string
	apiResourcePath string
	apiUser         string
	apiKey          string
	contentType     string
	adminUser       string
	adminPwd        string

	grantType = "password"
	scope     = "email profile pydio"
)

func init() {

	// TODO make this available via config. This only work on a local workstation for the time being.
	hostURL = "http://192.168.0.140:8080"
	apiResourcePath = "/a"

	apiUser = "cells-front"
	apiKey = "bLYmNHYUhiBMu5tBUUDrlSFM"
	contentType = "application/x-www-form-urlencoded"

	adminUser = "admin"
	adminPwd = "an admin password"
}

func GetPreparedApiClient() (*sdkapi.APIClient, context.Context, error) {
	config := sdkapi.NewConfiguration()
	config.BasePath = hostURL + apiResourcePath
	apiClient := sdkapi.NewAPIClient(config)

	ctx, err := withAuth(context.Background())
	if err != nil {
		return apiClient, nil, err
	}

	return apiClient, ctx, nil
}

func withAuth(ctx context.Context) (context.Context, error) {

	jwt, err := retrieveToken(adminUser, adminPwd)
	if err != nil {
		fmt.Printf("could not retrieve token: %s\n", err.Error())
		return nil, err
	}

	return context.WithValue(ctx, sdkapi.ContextAccessToken, jwt), nil
}

func retrieveToken(login, pwd string) (string, error) {

	fullURL := hostURL + "/auth/dex/token"

	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Add("username", login)
	data.Add("password", pwd)
	data.Add("scope", scope)
	data.Add("nonce", "aVerySpecialNonce") // TODO enhance this

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // Important our dex API does not yet support json payload.
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", getBasicAuthToken(apiUser, apiKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var respMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&respMap)
	if err != nil {
		return "", err
	}

	return respMap["id_token"].(string), nil
}

func getBasicAuthToken(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
