package s3

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/logging"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
	http2 "github.com/pydio/cells-sdk-go/v5/transport/http"
)

// LoadAwsConfig prepares a valid S3 configuration to create a new S3 client.
func LoadAwsConfig(ctx context.Context, store transport.ConfigStore, sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (aws.Config, error) {

	s3CredProv, err := NewCredentialsProvider(store, sdc)
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

	var httpClient *http.Client
	if s3c.RequestTimout > 0 {
		// cells_sdk.Log("Using transport with custom timeout of %d seconds", s3c.RequestTimout)
		httpClient = http2.GetClientWithTimeout(sdc, time.Duration(s3c.RequestTimout)*time.Second)
	} else {
		httpClient = http2.GetClient(sdc)
	}

	return config.LoadDefaultConfig(
		ctx,
		config.WithRegion(s3c.Region),
		config.WithHTTPClient(httpClient),
		config.WithCredentialsProvider(s3CredProv),
	)
}

// WithLogger enriches an AWS configuration with a logger that uses the passed writter and logs the defined S3 events.
func WithLogger(cfg aws.Config, writer io.Writer, logMode aws.ClientLogMode) aws.Config {
	cfg.Logger = logging.NewStandardLogger(writer)
	cfg.ClientLogMode = logMode
	// e.G: aws.LogSigning | aws.LogRetries | aws.LogRequest | aws.LogResponse | aws.LogDeprecatedUsage | aws.LogRequestEventMessage | aws.LogResponseEventMessage
	return cfg
}

// NewClientFromConfig uses the passed configuration and endpoint to create a fully configured S3 client.
func NewClientFromConfig(cfg aws.Config, endpoint string) *s3.Client {

	fo := func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		//	o.S3DisableContentMD5Validation = aws.Bool(true)
	}
	return s3.NewFromConfig(cfg, fo)
}

// func GetClient(store transport.ConfigStore, sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.Client, error) {

// 	s3CredProv, err := NewCredentialsProvider(store, sdc)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	if sdc.UseTokenCache {
// 		s3CredProv = aws.NewCredentialsCache(s3CredProv, func(options *aws.CredentialsCacheOptions) {
// 			// configuration of credential expiry window and jitter.
// 			options.ExpiryWindow = 10
// 			options.ExpiryWindowJitterFrac = 0.2
// 		})
// 	}

// 	var httpClient *http.Client
// 	if s3c.RequestTimout > 0 {
// 		// cells_sdk.Log("Using transport with custom timeout of %d seconds", s3c.RequestTimout)
// 		httpClient = http2.GetClientWithTimeout(sdc, time.Duration(s3c.RequestTimout)*time.Second)
// 	} else {
// 		httpClient = http2.GetClient(sdc)
// 	}

// 	cfg, err := config.LoadDefaultConfig(
// 		context.TODO(),
// 		config.WithRegion(s3c.Region),
// 		config.WithHTTPClient(httpClient),
// 		config.WithCredentialsProvider(s3CredProv),
// 	)
// 	if err != nil {
// 		log.Fatal("cannot load default S3 session configuration:", err.Error())
// 	}

// 	logWriter := printLnWriter{}

// 	logWriter.Println("... Got a logger")
// 	cfg.Logger = logging.NewStandardLogger(logWriter)
// 	cfg.ClientLogMode = aws.LogSigning | aws.LogRetries | aws.LogRequest | aws.LogResponse | aws.LogDeprecatedUsage | aws.LogRequestEventMessage | aws.LogResponseEventMessage

// 	fo := func(o *s3.Options) {
// 		o.BaseEndpoint = aws.String(s3c.Endpoint)
// 		//	o.S3DisableContentMD5Validation = aws.Bool(true)
// 	}

// 	return s3.NewFromConfig(cfg, fo), nil
// }
