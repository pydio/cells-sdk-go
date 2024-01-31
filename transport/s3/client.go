package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	http2 "github.com/pydio/cells-sdk-go/v5/transport"
	"log"
)

// LoadConfig prepares a valid S3 configuration to create a new S3 client.
func LoadConfig(ctx context.Context, sdc *cells_sdk.SdkConfig, options ...interface{}) (aws.Config, error) {

	s3CredProv, err := NewCredentialsProvider(sdc)
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

	// apply defined aws config options, e.G to set a custom region.
	for _, o := range options {
		switch typed := o.(type) {
		case cells_sdk.AwsConfigOption:
			cfg = typed(cfg)
		}
	}
	// cfg.Region = s3c.Region

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
