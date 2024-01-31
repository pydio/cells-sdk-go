package transport

import "github.com/pydio/cells-sdk-go/v5"

type StaticTokenProvider struct {
	staticToken string
}

func NewStaticTokenProvider(token string) cells_sdk.TokenProvider {
	return &StaticTokenProvider{staticToken: token}
}

func (t *StaticTokenProvider) Retrieve() (string, error) {
	return t.staticToken, nil
}

func (t *StaticTokenProvider) Expired() bool {
	return false
}
