package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

type LegacyCredentialsProvider struct {
	config *cellssdk.SdkConfig
	store  cellssdk.ConfigRefresher
}

func (lcp *LegacyCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

	if lcp.config.User == "" || lcp.config.Password == "" {
		return aws.Credentials{}, fmt.Errorf("cannot connect without login and password")
	}

	tp, e := transport.NewLegacyTokenProvider(lcp.config)
	if e != nil {
		return aws.Credentials{}, e
	}
	token, e := tp.Retrieve(ctx)
	if e != nil {
		return aws.Credentials{}, fmt.Errorf("could not retrieve token from legacy credentials, cause: %s", e.Error())
	}
	cred := aws.Credentials{
		AccessKeyID:     token,
		SecretAccessKey: cellssdk.DefaultS3ApiSecret,
		SessionToken:    "", // TODO
		Source:          lcp.config.Url,
		CanExpire:       true,
		Expires:         tp.ExpiresAt(),
	}
	return cred, nil
}
