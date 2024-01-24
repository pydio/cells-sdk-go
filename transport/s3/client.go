package s3

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
	http2 "github.com/pydio/cells-sdk-go/v5/transport/http"
)

func GetClient(store transport.ConfigStore, sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.Client, error) {

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

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(s3c.Region),
		config.WithHTTPClient(httpClient),
		config.WithCredentialsProvider(s3CredProv),
	)

	if err != nil {
		log.Fatal("cannot load default S3 session configuration:", err.Error())
	}

	fo := func(o *s3.Options) {
		o.BaseEndpoint = aws.String(s3c.Endpoint)
		//	o.S3DisableContentMD5Validation = aws.Bool(true)
	}

	return s3.NewFromConfig(cfg, fo), nil
}
