package config

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

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

func retrieveToken(sdkConfig *SdkConfig) (string, error) {

	cached := store.TokenFor(sdkConfig)
	if cached != "" {
		fmt.Println("[Auth: Retrieved token from cache]")
		return cached, nil
	}
	fmt.Println("[Auth: Loading Auth Token]")

	fullURL := sdkConfig.Protocol + "://" + sdkConfig.Url + OidcResourcePath + "/token"

	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Add("username", sdkConfig.User)
	data.Add("password", sdkConfig.Password)
	data.Add("scope", scope)
	data.Add("nonce", "aVerySpecialNonce")

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // Important our dex API does not yet support json payload.
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", basicAuthToken(sdkConfig.ClientKey, sdkConfig.ClientSecret))

	res, err := getHttpClient(sdkConfig).Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var respMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&respMap)
	if err != nil {
		return "", err
	}
	if errMsg, exists := respMap["error"]; exists {
		return "", fmt.Errorf("could not retrieve token, %s: %s", errMsg, respMap["error_description"])
	}

	// for k, v := range respMap {
	// 	fmt.Printf("%s - %v\n", k, v)
	// }

	token := respMap["id_token"].(string)

	expiry := respMap["expires_in"].(float64) - 60 // Secure by shortening expiration time
	store.Store(sdkConfig, token, time.Duration(expiry)*time.Second)
	return token, nil
}

func getHttpClient(sdkConfig *SdkConfig) *http.Client {

	if sdkConfig.SkipVerify {
		log.Println("[WARNING] Using SkipVerify for ignoring SSL certificate issues!!")
		return &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}}
	}
	return http.DefaultClient
}

func basicAuthToken(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
