/*
 * Copyright (c) 2019. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package oidc

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	cache "github.com/patrickmn/go-cache"

	cells_sdk "github.com/pydio/cells-sdk-go"
	"github.com/pydio/cells-sdk-go/client"
	"github.com/pydio/cells-sdk-go/client/frontend_service"
	"github.com/pydio/cells-sdk-go/models"
)

var (
	store = NewTokenStore()
)

type TokenStore struct {
	internalCache *cache.Cache
}

func NewTokenStore() *TokenStore {
	t := &TokenStore{
		internalCache: cache.New(20*time.Minute, 10*time.Minute),
	}
	return t
}

func (t *TokenStore) Store(c *cells_sdk.SdkConfig, token string, expiry time.Duration) {
	//fmt.Println("[Auth] Storing token with expiration ", expiry)
	t.internalCache.Set(t.computeKey(c), token, expiry)
}

func (t *TokenStore) TokenFor(c *cells_sdk.SdkConfig) string {

	if token, ok := t.internalCache.Get(t.computeKey(c)); ok {
		return token.(string)
	}
	return ""

}

func (t *TokenStore) computeKey(c *cells_sdk.SdkConfig) string {
	// Is this relly necessary or rather security theater?
	// using a generic password causes issues when testing wrong password access.
	//s := fmt.Sprintf("%s-%s-%s-%s-%s", c.Url, c.ClientKey, c.ClientSecret, c.User, "OBFUSCATED PWD XXXX")
	var s string
	if c.IdToken != "" {
		s = fmt.Sprintf("%s-%s", c.Url, c.User)
	} else {
		s = fmt.Sprintf("%s-%s-%s-%s-%s", c.Url, c.ClientKey, c.ClientSecret, c.User, c.Password)
	}
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RetrieveToken(sdkConfig *cells_sdk.SdkConfig) (string, error) {

	if cached := store.TokenFor(sdkConfig); sdkConfig.UseTokenCache && cached != "" {
		return cached, nil
	}

	if sdkConfig.IdToken != "" {
		// We passed a pre-fetched valid token
		expTime := time.Unix(int64(sdkConfig.TokenExpiresAt), 0)
		store.Store(sdkConfig, sdkConfig.IdToken, expTime.Sub(time.Now()))
		return sdkConfig.IdToken, nil
	}

	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return "", e
	}

	ctx := context.Background()
	runtime := httptransport.New(u.Host, "/a", []string{u.Scheme})
	if sdkConfig.SkipVerify {
		runtime.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired and self-signed SSL certificates
		}
	}
	if sdkConfig.CustomHeaders != nil && len(sdkConfig.CustomHeaders) > 0 {
		runtime.Transport = &customHeaderRoundTripper{
			rt:      runtime.Transport,
			Headers: sdkConfig.CustomHeaders,
		}
	}
	runtime.Context = ctx

	frontRequest := frontend_service.NewFrontSessionParams().WithBody(&models.RestFrontSessionRequest{
		AuthInfo: map[string]string{
			"login":    sdkConfig.User,
			"password": sdkConfig.Password,
			"type":     "credentials",
		},
	}).WithContext(ctx)
	resp, err := client.New(runtime, strfmt.Default).FrontendService.FrontSession(frontRequest)
	if err != nil {
		return "", err
	}
	token := resp.Payload.JWT
	expiryDate := time.Unix(int64(resp.Payload.ExpireTime), 0).Add(-60 * time.Second)

	store.Store(sdkConfig, token, expiryDate.Sub(time.Now()))
	return token, nil
}

func basicAuthToken(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

type customHeaderRoundTripper struct {
	rt      http.RoundTripper
	Headers map[string]string
}

func (c customHeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	return c.rt.RoundTrip(req)
}
