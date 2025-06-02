package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
)

type OAuthCredentialsProvider struct {
	config *apiv1.SdkConfig
	store  apiv1.ConfigRefresher
}

func (ocp *OAuthCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	none := aws.Credentials{}
	if ocp.store == nil {
		return none, fmt.Errorf("cannot retrieve OAuth credentials without a store")
	}
	_, err := ocp.store.RefreshIfRequired(ctx, ocp.config)
	if err != nil {
		return none, err
	}
	if ocp.config.IdToken == "" {
		return aws.Credentials{}, fmt.Errorf("id token cannot be empty at this stage")
	}

	// We generally refresh the token 60 seconds before its expiration
	// ==> underlying problem: call via AWS SDK fail when the token is refreshed in another thread
	//     because the current idToken then becomes invalid.
	//  This is not very resilient and must be improved. TODO
	expiration := time.Unix(int64(ocp.config.TokenExpiresAt), 0).Add(-65 * time.Second)
	currCredentials := aws.Credentials{
		AccessKeyID:     ocp.config.IdToken,
		SecretAccessKey: apiv1.DefaultS3ApiSecret,
		SessionToken:    "", // TODO
		Source:          ocp.config.Url,
		CanExpire:       true,
		Expires:         expiration,
	}
	return currCredentials, nil
}

func (ocp *OAuthCredentialsProvider) SetConfigRefresher(store apiv1.ConfigRefresher) {
	ocp.store = store
}
