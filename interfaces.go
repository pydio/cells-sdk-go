package cells_sdk

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"net/http"
)

type ConfigStore interface {
	RefreshIfRequired(*SdkConfig) (bool, error)
}

type CellsCredentialsProvider interface {
	// Retrieve returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve(ctx context.Context) (aws.Credentials, error)
	SetConfigStore(store ConfigStore)
}

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}

type Option func(t *http.Transport) *http.Transport

type RoundTripOption func(t http.RoundTripper) http.RoundTripper

type HttpClientOption func(client *http.Client) *http.Client

type AwsConfigOption func(config aws.Config) aws.Config

type CredentialProviderOption func(provider CellsCredentialsProvider) CellsCredentialsProvider
