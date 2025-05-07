package s3

import (
	"context"
	"fmt"
	"github.com/pydio/cells-sdk-go/v4/apiv1"
	http2 "github.com/pydio/cells-sdk-go/v4/apiv1/transport"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// LoadConfig prepares a valid S3 configuration to create a new S3 client.
// It returns an error if the passed option are not one of the supported type that are:
// cellsSdk.HttpClientOption, cellsSdk.TransportOption, cellsSdk.RoundTripOption, cellsSdk.AwsConfigOption, cellsSdk.CredentialProviderOption
func LoadConfig(ctx context.Context, sdc *apiv1.SdkConfig, options ...interface{}) (aws.Config, error) {

	for _, o := range options {
		switch o.(type) {
		case apiv1.HttpClientOption, apiv1.TransportOption, apiv1.RoundTripOption, apiv1.AwsConfigOption, apiv1.CredentialProviderOption:
		default:
			return aws.Config{}, fmt.Errorf("unsupported option type")
		}
	}

	s3CredProv, err := NewCredentialsProvider(sdc, options...)
	if err != nil {
		log.Fatal(err.Error())
	}

	if sdc.UseTokenCache {
		s3CredProv = aws.NewCredentialsCache(s3CredProv, func(options *aws.CredentialsCacheOptions) {
			// configuration of credential expiry window and jitter.
			options.ExpiryWindow = 10
			options.ExpiryWindowJitterFrac = 0.2
		})
	}

	httpClient := http2.NewHttpClient(sdc, options...)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithHTTPClient(httpClient),
		config.WithCredentialsProvider(s3CredProv),
	)

	// Apply defined AWS config options, e.G. to set a custom region.
	for _, o := range options {
		switch typed := o.(type) {
		case apiv1.AwsConfigOption:
			cfg = typed(cfg)
		}
	}

	// This can be overwritten by an AwsConfigOption that are applied just above
	if cfg.Region == "" {
		cfg.Region = apiv1.DefaultS3Region
	}

	return cfg, err
}

// NewClientFromConfig is a helper function to create a fully configured S3 client easily.
// It sets the passed endpoint and wraps the call to AWS SDK S3 default method.
func NewClientFromConfig(cfg aws.Config, endpoint string) *s3.Client {

	fo := func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		//	o.S3DisableContentMD5Validation = aws.Bool(true)
	}
	return s3.NewFromConfig(cfg, fo)
}
