package cells_sdk

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ConfigStore interface {
	RefreshIfRequired(context.Context, *SdkConfig) (bool, error)
}

type CellsCredentialsProvider interface {

	// Retrieve returns valid AWS credentials or an error
	// if the value were not obtainable, or empty.
	Retrieve(context.Context) (aws.Credentials, error)
	SetConfigStore(ConfigStore)
}

type TokenProvider interface {
	Retrieve(context.Context) (string, error)
	Expired() bool
}

type Option func(*http.Transport) *http.Transport

type RoundTripOption func(http.RoundTripper) http.RoundTripper

type HttpClientOption func(*http.Client) *http.Client

type AwsConfigOption func(aws.Config) aws.Config

type CredentialProviderOption func(CellsCredentialsProvider) CellsCredentialsProvider
