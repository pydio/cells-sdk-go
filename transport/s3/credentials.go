package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

func NewCredentialsProvider(sdc *cellssdk.SdkConfig, options ...interface{}) (aws.CredentialsProvider, error) {

	var provider cellssdk.CellsCredentialsProvider

	switch sdc.AuthType {
	case cellssdk.AuthTypeOAuth:
		provider = &OAuthCredentialsProvider{config: sdc}
	case cellssdk.AuthTypePat:
		provider = &PatCredentialsProvider{config: sdc}
	case cellssdk.AuthTypeClientAuth:
		provider = &LegacyCredentialsProvider{config: sdc}
	default:
		return nil, fmt.Errorf("unsupported auth type %s, we cannot create a relevant AWS provider", sdc.AuthType)
	}

	// Apply passed options if any of them is relevant here for this class, e.G to set CellsConfigStore.
	for _, o := range options {
		switch typed := o.(type) {
		case cellssdk.CredentialProviderOption:
			provider = typed(provider)
		}
	}

	return provider, nil
}

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

func (pcp *PatCredentialsProvider) SetConfigStore(_ cellssdk.ConfigStore) {}

type LegacyCredentialsProvider struct {
	config *cellssdk.SdkConfig
	store  cellssdk.ConfigStore
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

func (lcp *LegacyCredentialsProvider) SetConfigStore(_ cellssdk.ConfigStore) {}

type OAuthCredentialsProvider struct {
	config *cellssdk.SdkConfig
	store  cellssdk.ConfigStore
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

func (ocp *OAuthCredentialsProvider) SetConfigStore(store cellssdk.ConfigStore) {
	ocp.store = store
}
