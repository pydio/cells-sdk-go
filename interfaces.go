package cells_sdk

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"net/http"
)

type ConfigStore interface {
	RefreshIfRequired(*SdkConfig) (bool, error)
}

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}

type Option func(t *http.Transport) *http.Transport

type RoundTripOption func(t http.RoundTripper) http.RoundTripper

type HttpClientOption func(client *http.Client) *http.Client

type AwsConfigOption func(config aws.Config) aws.Config
