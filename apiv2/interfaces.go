package cells_sdk

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type TokenProvider interface {
	Retrieve(context.Context) (string, error)
	Expired() bool
	ExpiresAt() time.Time
}

type ConfigRefresher interface {
	RefreshIfRequired(context.Context, *SdkConfig) (bool, error)
}

type ConfigRefresherConsumer interface {
	SetConfigRefresher(ConfigRefresher)
}

type HttpClientOption func(*http.Client) *http.Client

type TransportOption func(*http.Transport) *http.Transport

type RoundTripOption func(http.RoundTripper) http.RoundTripper

type AwsConfigOption func(aws.Config) aws.Config

type CredentialProviderOption func(aws.CredentialsProvider) aws.CredentialsProvider
