package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

func NewCredentialsProvider(store transport.ConfigStore, sdc *cells_sdk.SdkConfig) (aws.CredentialsProvider, error) {
	switch sdc.AuthType {
	case cells_sdk.AuthTypeOAuth:
		return &OAuthCredentialsProvider{store: store, config: sdc}, nil
	case cells_sdk.AuthTypePat:
		return &PatCredentialsProvider{sdc}, nil
	case cells_sdk.AuthTypeClientAuth:
		return &LegacyCredentialsProvider{sdc}, nil
	default:
		return nil, fmt.Errorf("unsupported auth type %s, we cannot create a relevant AWS provider", sdc.AuthType)
	}

}

type PatCredentialsProvider struct {
	config *cells_sdk.SdkConfig
}

func (pcp *PatCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

	if pcp.config.IdToken == "" {
		return aws.Credentials{}, fmt.Errorf("cannot provide credentials with an empty PAT")
	}

	return aws.Credentials{
		AccessKeyID:     pcp.config.IdToken,
		SecretAccessKey: transport.CellsS3SecretDefault,
		SessionToken:    "",
		Source:          pcp.config.Url,
		CanExpire:       false,
	}, nil
}

type LegacyCredentialsProvider struct {
	config *cells_sdk.SdkConfig
}

func (lcp *LegacyCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

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
		SecretAccessKey: transport.CellsS3SecretDefault,
		SessionToken:    "", // TODO
		Source:          lcp.config.Url,
		CanExpire:       true,
		Expires:         tp.ExpiresAt(),
	}
	return cred, nil
}

type OAuthCredentialsProvider struct {
	store  transport.ConfigStore
	config *cells_sdk.SdkConfig
}

func (ocp *OAuthCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

	expiration := time.Unix(int64(ocp.config.TokenExpiresAt), 0)

	currCreds := aws.Credentials{
		AccessKeyID:     ocp.config.IdToken,
		SecretAccessKey: transport.CellsS3SecretDefault,
		SessionToken:    "", // TODO
		Source:          ocp.config.Url,
		CanExpire:       true,
		Expires:         expiration,
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
