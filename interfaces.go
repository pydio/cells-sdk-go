package cells_sdk

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type TokenProvider interface {
	Retrieve(context.Context) (string, error)
	Expired() bool
}

type CellsCredentialsProvider interface {
	aws.CredentialsProvider
	SetConfigStore(ConfigStore)
}

type ConfigStore interface {
	RefreshIfRequired(context.Context, *SdkConfig) (bool, error)
}

type Option func(*http.Transport) *http.Transport

type RoundTripOption func(http.RoundTripper) http.RoundTripper

type HttpClientOption func(*http.Client) *http.Client

type AwsConfigOption func(aws.Config) aws.Config

type CredentialProviderOption func(CellsCredentialsProvider) CellsCredentialsProvider
