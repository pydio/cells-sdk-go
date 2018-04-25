package util

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	// TODO make this available via config. This only work on a local workstation for the time being.
	hostURL     = "http://192.168.122.89:8080"
	apiUser     = "pydio-frontend"
	apiKey      = "hj8tZI5Cca2YPrAWL17ZUQJH"
	contentType = "application/x-www-form-urlencoded"

	grantType = "password"
	scope     = "email profile pydio"
)

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
