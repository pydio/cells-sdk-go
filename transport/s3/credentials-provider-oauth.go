package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

type OAuthCredentialsProvider struct {
	config *cellssdk.SdkConfig
	store  cellssdk.ConfigRefresher
}

func (ocp *OAuthCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

	expiration := time.Unix(int64(ocp.config.TokenExpiresAt), 0)

	currCredentials := aws.Credentials{
		AccessKeyID:     ocp.config.IdToken,
		SecretAccessKey: cellssdk.DefaultS3ApiSecret,
		SessionToken:    "", // TODO
		Source:          ocp.config.Url,
		CanExpire:       true,
		Expires:         expiration,
	}

	if ocp.store == nil {
		return currCredentials, fmt.Errorf("cannot retrieve OAuth credentials without a store")
	}
	refreshed, err := ocp.store.RefreshIfRequired(ctx, ocp.config)
	if err != nil {
		return aws.Credentials{}, err
	}
	if !refreshed {
		return currCredentials, nil
	}

	currCredentials.AccessKeyID = ocp.config.IdToken
	currCredentials.Expires = time.Unix(int64(ocp.config.TokenExpiresAt), 0)
	return currCredentials, nil
}

func (ocp *OAuthCredentialsProvider) SetConfigRefresher(store cellssdk.ConfigRefresher) {
	ocp.store = store
}
