package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

func NewCredentialsProvider(sdc *cells_sdk.SdkConfig, options ...interface{}) (aws.CredentialsProvider, error) {

	var provider cells_sdk.CellsCredentialsProvider

	switch sdc.AuthType {
	case cells_sdk.AuthTypeOAuth:
		provider = &OAuthCredentialsProvider{config: sdc}
		break
	case cells_sdk.AuthTypePat:
		provider = &PatCredentialsProvider{config: sdc}
		break
	case cells_sdk.AuthTypeClientAuth:
		provider = &LegacyCredentialsProvider{config: sdc}
		break
	default:
		return nil, fmt.Errorf("unsupported auth type %s, we cannot create a relevant AWS provider", sdc.AuthType)
	}

	// Apply passed options if any of them is relevant here for this class, e.G to set CellsConfigStore.
	for _, o := range options {
		switch typed := o.(type) {
		case cells_sdk.CredentialProviderOption:
			provider = typed(provider)
			// 	break
			// default:
			// 	fmt.Println("... NOT a Cred prov option")
		}
	}

	return provider, nil
}

type PatCredentialsProvider struct {
	config *cells_sdk.SdkConfig
}

func (pcp *PatCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {

	if pcp.config.IdToken == "" {
		return aws.Credentials{}, fmt.Errorf("cannot provide credentials with an empty PAT")
	}

	return aws.Credentials{
		AccessKeyID:     pcp.config.IdToken,
		SecretAccessKey: cells_sdk.DefaultS3ApiSecret,
		SessionToken:    "",
		Source:          pcp.config.Url,
		CanExpire:       false,
	}, nil
}

func (pcp *PatCredentialsProvider) SetConfigStore(store cells_sdk.ConfigStore) {}

type LegacyCredentialsProvider struct {
	config *cells_sdk.SdkConfig
	store  cells_sdk.ConfigStore
}

func (lcp *LegacyCredentialsProvider) Retrieve(_ context.Context) (aws.Credentials, error) {

	if lcp.config.User == "" || lcp.config.Password == "" {
		return aws.Credentials{}, fmt.Errorf("cannot connect without login and password")
	}

	tp, e := transport.NewLegacyTokenProvider(lcp.config)
	if e != nil {
		return aws.Credentials{}, e
	}
	token, e := tp.Retrieve()
	if e != nil {
		return aws.Credentials{}, fmt.Errorf("could not retrieve token from legacy credentials, cause: %s", e.Error())
	}
	cred := aws.Credentials{
		AccessKeyID:     token,
		SecretAccessKey: cells_sdk.DefaultS3ApiSecret,
		SessionToken:    "", // TODO
		Source:          lcp.config.Url,
		CanExpire:       true,
		Expires:         tp.ExpiresAt(),
	}
	return cred, nil
}

func (pcp *LegacyCredentialsProvider) SetConfigStore(store cells_sdk.ConfigStore) {}

type OAuthCredentialsProvider struct {
	config *cells_sdk.SdkConfig
	store  cells_sdk.ConfigStore
}

func (ocp *OAuthCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

	expiration := time.Unix(int64(ocp.config.TokenExpiresAt), 0)

	currCreds := aws.Credentials{
		AccessKeyID:     ocp.config.IdToken,
		SecretAccessKey: cells_sdk.DefaultS3ApiSecret,
		SessionToken:    "", // TODO
		Source:          ocp.config.Url,
		CanExpire:       true,
		Expires:         expiration,
	}

	if ocp.store == nil {
		return currCreds, fmt.Errorf("cannot retrieve OAuth credentials without a store")
	}
	refreshed, err := ocp.store.RefreshIfRequired(ocp.config)
	if err != nil {
		return aws.Credentials{}, err
	}
	if !refreshed {
		return currCreds, nil
	}

	currCreds.AccessKeyID = ocp.config.IdToken
	currCreds.Expires = time.Unix(int64(ocp.config.TokenExpiresAt), 0)
	return currCreds, nil
}

func (ocp *OAuthCredentialsProvider) SetConfigStore(store cells_sdk.ConfigStore) {
	ocp.store = store
}
