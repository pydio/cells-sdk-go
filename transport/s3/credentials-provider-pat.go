package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

type PatCredentialsProvider struct {
	config *cellssdk.SdkConfig
}

func (pcp *PatCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {

	if pcp.config.IdToken == "" {
		return aws.Credentials{}, fmt.Errorf("cannot provide credentials with an empty PAT")
	}

	return aws.Credentials{
		AccessKeyID:     pcp.config.IdToken,
		SecretAccessKey: cellssdk.DefaultS3ApiSecret,
		SessionToken:    "",
		Source:          pcp.config.Url,
		CanExpire:       false,
	}, nil
}
