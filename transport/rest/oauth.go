package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	cellshttp "github.com/pydio/cells-sdk-go/v5/transport"
)

const (
	// DefaultCallbackUrl is the default URL that is registered for cells-client in the Server out of the box.
	DefaultCallbackUrl = "http://localhost:3000/servers/callback"
	// NoBrowserCallbackSuffix is the Uri that is used to display a code that can be copy/pasted by the end user.
	NoBrowserCallbackSuffix = "/oauth2/oob"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	// if the server returns an error, we will have this field so that we can check.
	StatusCode int `json:"status_code"`
}

// OAuthPrepareUrl creates a URL that can be opened in a browser or copy/pasted by the end user.
func OAuthPrepareUrl(clientId, state, serverUrl, callbackUrl string) (preparedUrl string, e error) {

	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("client_id", clientId)
	values.Add("scope", "openid email profile pydio offline")
	values.Add("state", state)
	values.Add("redirect_uri", callbackUrl)

	authU, e := url.Parse(serverUrl)
	if e != nil {
		return
	}
	authU.RawQuery = values.Encode()
	authU.Path = "/oidc/oauth2/auth"
	preparedUrl = authU.String()
	return
}

// OAuthExchangeCode retrieves an AccessToken/RefreshToken pair using the passed OAuth exchange code.
// It then updates the passed configuration.
func OAuthExchangeCode(c *cellssdk.SdkConfig, clientId, code, callbackUrl string) error {

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", callbackUrl)
	values.Add("client_id", clientId)

	tokenU, err := url.Parse(c.Url)
	if err != nil {
		return err
	}
	tokenU.Path = "/oidc/oauth2/token"

	httpClient := cellshttp.NewHttpClient(c)
	resp, err := httpClient.Post(tokenU.String(), "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r tokenResponse
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if r.StatusCode > 299 {
		return fmt.Errorf("could not perfom authentication flow: response body %s", string(b))
	}

	c.IdToken = r.AccessToken
	c.RefreshToken = r.RefreshToken
	c.TokenExpiresAt = int(time.Now().Unix()) + r.ExpiresIn

	return nil
}

func RefreshJwtToken(clientId string, sdkConfig *cellssdk.SdkConfig) (bool, error) {

	// Not yet expired, ignore
	if time.Unix(int64(sdkConfig.TokenExpiresAt), 0).After(time.Now().Add(60 * time.Second)) {
		return false, nil
	}
	data := url.Values{}
	data.Add("grant_type", "refresh_token")
	data.Add("client_id", clientId)
	data.Add("refresh_token", sdkConfig.RefreshToken)
	data.Add("scope", "openid email profile pydio offline")

	httpReq, err := http.NewRequest("POST", sdkConfig.Url+"/oidc/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return true, err
	}
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Cache-Control", "no-cache")

	client := cellshttp.NewHttpClient(sdkConfig)
	res, err := client.Do(httpReq)
	if err != nil {
		return false, err
	} else if res.StatusCode != 200 {
		bb, _ := io.ReadAll(res.Body)
		return false, fmt.Errorf("received status code %d - %s", res.StatusCode, string(bb))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Warning: could not close request body stream: %s", err.Error())
		}
	}(res.Body)
	var respMap tokenResponse
	err = json.NewDecoder(res.Body).Decode(&respMap)
	if err != nil {
		return true, fmt.Errorf("could not unmarshall response with status %d: %s\nerror cause: %s", res.StatusCode, res.Status, err.Error())
	}
	sdkConfig.IdToken = respMap.AccessToken
	sdkConfig.RefreshToken = respMap.RefreshToken
	sdkConfig.TokenExpiresAt = int(time.Now().Unix()) + respMap.ExpiresIn

	return true, nil
}
