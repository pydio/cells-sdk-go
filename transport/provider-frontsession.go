package transport

import (
	"context"
	"net/url"
	"time"

	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/client"
	"github.com/pydio/cells-sdk-go/v4/client/frontend_service"
	"github.com/pydio/cells-sdk-go/v4/models"
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

func NewFrontSessionTokenProvider(c *cells_sdk.SdkConfig) (TokenProvider, error) {
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

func (f *FrontSessionTokenProvider) Retrieve() (string, error) {
	if !f.Expired() {
		return f.token, nil
	}
	ctx := context.Background()
	runtime := openapi.New(f.url.Host, CellsApiResourcePath, []string{f.url.Scheme})
	runtime.Context = ctx
	runtime.Transport = New(
		WithSkipVerify(f.skipVerify),
		WithCustomHeaders(f.customHeaders),
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
	expiryDate := time.Unix(int64(resp.Payload.ExpireTime), 0).Add(-60 * time.Second)
	f.token = token
	f.expiryDate = expiryDate
	return f.token, nil
}

func (f *FrontSessionTokenProvider) Expired() bool {
	return f.token == "" || f.expiryDate.Before(time.Now())
}

func LoadAccessToken(c *cells_sdk.SdkConfig) (string, error) {
	p, e := NewFrontSessionTokenProvider(c)
	if e != nil {
		return "", e
	}
	return p.Retrieve()
}
