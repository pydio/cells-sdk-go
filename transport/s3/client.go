package s3

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/transport"
	http2 "github.com/pydio/cells-sdk-go/v4/transport/http"
)

func GetClient(store transport.ConfigStore, sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.Client, error) {

	s3CredProv := NewCredentialsProvider(store, sdc)

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

	// tmpCli := s3.NewFromConfig(cfg, {}options, )

	// options = append(optFns,
	// 	func(o *s3.Options) {
	// 		o.BaseEndpoint = aws.String(s3c.Endpoint),
	// 		o.S3DisableContentMD5Validation = aws.Bool(true),
	// 	},
	// 	// func(o *s3.Options) { o.S3DisableContentMD5Validation = aws.Bool(true) },
	// )

	return s3.NewFromConfig(cfg, fo), nil
}

// // GetClient creates and configure a new S3 client at each request.
// func GetClient(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.S3, error) {

// 	tp, e := transport.TokenProviderFromConfig(sdc)
// 	if e != nil {
// 		return nil, e
// 	}
// 	htCl := http2.GetClient(sdc)
// 	conf := aws.NewConfig()
// 	conf.WithCredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp)))
// 	conf.WithHTTPClient(htCl)
// 	conf.WithEndpoint(s3c.Endpoint)
// 	conf.WithRegion(s3c.Region)
// 	s3Session, err := session.NewSession(conf)
// 	if err != nil {
// 		log.Fatal("cannot initialise default S3 session: " + err.Error())
// 	}
// 	return s3.New(s3Session), nil

// }

// // LoadS3Config retrieves current S3 configuration to start a new client
// func LoadS3Config(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (aws.Config, error) {
// 	// tp, e := transport.TokenProviderFromConfig(sdc)
// 	// if e != nil {
// 	// 	return nil, e
// 	// }
// 	// htCl := http2.GetClient(sdc)

// 	// conf := aws.NewContransportfig()
// 	// conf.WithCredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp)))
// 	// conf.WithHTTPClient(htCl)
// 	// conf.WithEndpoint(s3c.Endpoint)
// 	// conf.WithRegion(s3c.Region)

// 	// s3Session, err := session.NewSession(conf)
// 	// if err != nil {
// 	// 	log.Fatal("cannot initialise default S3 session: " + err.Error())
// 	// }

// 	htCl := http2.GetClient(sdc)
// 	cfg, err := config.LoadDefaultConfig(
// 		context.TODO(),
// 		config.WithRegion(s3c.Region),
// 		// config.WithCtransportredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp))),
// 		config.WithHTTPClient(htCl),
// 		// config.WithEndpoint(s3c.Endpoint),
// 	)
// 	if err != nil {
// 		log.Fatal("cannot load default S3 session configuration:", err.Error())
// 	}
// 	return cfg, nil

// }
