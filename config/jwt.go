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

	"github.com/pydio/cells-sdk-go/api"
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
func GetPreparedApiClient(sdkConfig *SdkConfig) (*api.APIClient, context.Context, error) {
	apiConf := api.NewConfiguration()
	apiConf.HTTPClient = getHttpClient(sdkConfig)
	apiConf.BasePath = strings.TrimRight(sdkConfig.Url, "/") + ApiResourcePath
	apiClient := api.NewAPIClient(apiConf)

	ctx, err := withAuth(context.Background(), sdkConfig)
	if err != nil {
		return apiClient, nil, err
	}

	return apiClient, ctx, nil
}

func getHttpClient(sdkConfig *SdkConfig) *http.Client {

	if sdkConfig.SkipVerify {
		log.Println("[WARNING] Using SkipVerify for ignoring SSL certificate issues!!")
		return &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}}
	} else {
		return http.DefaultClient
	}

}

func withAuth(ctx context.Context, sdkConfig *SdkConfig) (context.Context, error) {

	jwt, err := retrieveToken(sdkConfig)
	if err != nil {
		fmt.Printf("could not retrieve token: %s\n", err.Error())
		return nil, err
	}

	return context.WithValue(ctx, api.ContextAccessToken, jwt), nil
}

func retrieveToken(sdkConfig *SdkConfig) (string, error) {

	cached := store.TokenFor(sdkConfig)
	if cached != "" {
		fmt.Println("[Auth: Retrieved token from cache]")
		return cached, nil
	}
	fmt.Println("[Auth: Loading Auth Token]")

	fullURL := sdkConfig.Url + OidcResourcePath + "/token"

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

	token := respMap["id_token"].(string)
	expiry := respMap["expires_in"].(float64) - 60 // Secure by shortening expiration time
	store.Store(sdkConfig, token, time.Duration(expiry)*time.Second)
	return token, nil
}

func basicAuthToken(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
