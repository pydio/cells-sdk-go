package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	cells_sdk "github.com/pydio/cells-sdk-go"
	http2 "github.com/pydio/cells-sdk-go/transport/http"
	"github.com/pydio/cells-sdk-go/transport/oidc"
)

var (
	apiResourcePath = "/a"
)

func GetRestClientTransport(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (context.Context, runtime.ClientTransport, error) {
	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, nil, e
	}

	ctx := context.Background()
	transport := httptransport.New(u.Host, apiResourcePath, []string{u.Scheme})
	if sdkConfig.SkipVerify {
		transport.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired and self-signed SSL certificates
		}
	}
	transport.Context = ctx

	if anonymous {
		return ctx, transport, nil
	}

	jwt, err := oidc.RetrieveToken(sdkConfig)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"cannot retrieve token with config:\n%s - %s - %s - %s - %s - %v\nerror cause: %s",
			sdkConfig.Url, sdkConfig.ClientKey, sdkConfig.ClientSecret,
			sdkConfig.User, sdkConfig.Password, sdkConfig.SkipVerify, err.Error())
	}
	bearerTokenAuth := httptransport.BearerToken(jwt)
	transport.DefaultAuthentication = bearerTokenAuth

	return ctx, transport, nil
}

// PrepareSimpleRequest returns a valid http client and pre-set request with headers.
func PrepareSimpleRequest(sdkConfig *cells_sdk.SdkConfig) (*http.Client, *http.Request, error) {
	h := make(http.Header)
	jwt, err := oidc.RetrieveToken(sdkConfig)
	if err != nil {
		return nil, nil, err
	}
	h.Set("Authorization", "Bearer "+jwt)
	request := &http.Request{
		Header: h,
	}
	return http2.GetHttpClient(sdkConfig), request, nil
}
