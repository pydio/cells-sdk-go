package transport

import (
	"context"
	"net/url"
	"time"

	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/client"
	"github.com/pydio/cells-sdk-go/v5/client/frontend_service"
	"github.com/pydio/cells-sdk-go/v5/models"
	"github.com/pydio/cells-sdk-go/v5/transport/http"
)

type FrontSessionTokenProvider struct {
	url           *url.URL
	user          string
	password      string
	skipVerify    bool
	customHeaders map[string]string
	token         string
	expiryDate    time.Time
}

func NewFrontSessionTokenProvider(c *cellssdk.SdkConfig) (cellssdk.TokenProvider, error) {
	u, e := url.Parse(c.Url)
	if e != nil {
		return nil, e
	}
	return &FrontSessionTokenProvider{
		url:           u,
		user:          c.User,
		password:      c.Password,
		skipVerify:    c.SkipVerify,
		customHeaders: c.CustomHeaders,
	}, nil
}

func NewLegacyTokenProvider(c *cellssdk.SdkConfig) (*FrontSessionTokenProvider, error) {
	u, e := url.Parse(c.Url)
	if e != nil {
		return nil, e
	}
	return &FrontSessionTokenProvider{
		url:           u,
		user:          c.User,
		password:      c.Password,
		skipVerify:    c.SkipVerify,
		customHeaders: c.CustomHeaders,
	}, nil
}

func (f *FrontSessionTokenProvider) Retrieve(ctx context.Context) (string, error) {
	if !f.Expired() {
		return f.token, nil
	}
	runtime := openapi.New(f.url.Host, CellsApiResourcePath, []string{f.url.Scheme})
	runtime.Context = ctx
	runtime.Transport = New(
		http.WithSkipVerify(f.skipVerify),
		http.WithCustomHeaders(f.customHeaders),
	)

	frontRequest := frontend_service.NewFrontSessionParams().WithBody(&models.RestFrontSessionRequest{
		AuthInfo: map[string]string{
			"login":    f.user,
			"password": f.password,
			"type":     "credentials",
		},
	}).WithContext(ctx)
	resp, err := client.New(runtime, strfmt.Default).FrontendService.FrontSession(frontRequest)
	if err != nil {
		return "", err
	}
	token := resp.Payload.JWT
	f.token = token
	f.expiryDate = time.Unix(int64(resp.Payload.ExpireTime), 0)
	return f.token, nil
}

func (f *FrontSessionTokenProvider) Expired() bool {
	return f.token == "" || f.expiryDate.Before(time.Now())
}

func (f *FrontSessionTokenProvider) ExpiresAt() time.Time {
	return f.expiryDate
}

//func LoadAccessToken(c *cells_sdk.SdkConfig) (string, error) {
//	p, e := NewFrontSessionTokenProvider(c)
//	if e != nil {
//		return "", e
//	}
//	return p.Retrieve(context.TODO())
//}
