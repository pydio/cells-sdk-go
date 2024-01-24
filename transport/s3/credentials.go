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
		return nil, fmt.Errorf("unsupported auth type %s, we cannot create a relevant AWS provider")
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

// type CellsCredentialsProvider struct {
// 	config *cells_sdk.SdkConfig
// }

// func (ccp *CellsCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {

// 	fmt.Println()
// 	fmt.Printf("... About to retrieve credentials for %s@%s\n", ccp.config.User, ccp.config.Url)

// 	if ccp.config.IdToken != "" {
// 		// TODO	handle expiration
// 		cred := aws.Credentials{
// 			AccessKeyID:     ccp.config.IdToken,
// 			SecretAccessKey: transport.CellsS3SecretDefault,
// 			SessionToken:    "", // TODO
// 			Source:          ccp.config.Url,
// 			CanExpire:       ccp.config.TokenExpiresAt > 0,
// 			Expires:         time.Unix(int64(ccp.config.TokenExpiresAt), 0).Add(-60 * time.Second),
// 		}

// 		fmt.Println()
// 		fmt.Printf("... Retrieved credentials with token for %s@%s\n", ccp.config.User, ccp.config.Url)

// 		return cred, nil

// 	} else if ccp.config.User != "" && ccp.config.Password != "" { // We check if we have a legacy login / pwd pair

// 		tp, e := transport.NewFrontSessionTokenProvider(ccp.config)
// 		if e != nil {
// 			return aws.Credentials{}, e
// 		}
// 		token, e := tp.Retrieve()
// 		if e != nil {
// 			fmt.Println("Could not retrieve token from legacy credentials")
// 			return aws.Credentials{}, e
// 		}
// 		cred := aws.Credentials{
// 			AccessKeyID:     token,
// 			SecretAccessKey: transport.CellsS3SecretDefault,
// 			SessionToken:    "", // TODO
// 			Source:          ccp.config.Url,
// 			CanExpire:       true,
// 			Expires:         time.Now().Add(300 * time.Second), // TODO (int64(ccp.config.TokenExpiresAt), 0).Add(-60 * time.Second)
// 		}
// 		return cred, nil

// 	} else {
// 		return aws.Credentials{}, nil
// 	}
// }

// type s3CredentialProvider struct {
// 	tp transport.TokenProvider
// }

// func TokenProviderFromConfig(c *cells_sdk.SdkConfig) (TokenProvider, error) {
// 	if c.IdToken != "" {
// 		return c, nil // SdkConfig implements TokenProvider interface
// 	} else {
// 		tp, e := transport.NewFrontSessionTokenProvider(c)
// 		if e != nil {
// 			return nil, e
// 		}
// 		if c.UseTokenCache {
// 			tp = WithProviderCache(tp, c)
// 		}
// 		return tp, nil
// 	}
// }

// type CredentialsProvider interface {
// 	// Retrieve returns nil if it successfully retrieved the value.
// 	// Error is returned if the value were not obtainable, or empty.
// 	Retrieve(ctx context.Context) (aws.Credentials, error)
// }

// func AsS3CredentialsProvider(provider transport.TokenProvider) aws.CredentialsProvider {
// 	return &s3CredentialProvider{
// 		tp: provider,
// 	}
// }

// // s3CredProv := AsS3CredentialsProvider(tp)

// func (s *s3CredentialProvider) Retrieve(_ context.Context) (aws.Credentials, error) {

// 	token, err := s.tp.Retrieve()
// 	if err != nil {
// 		log.Fatal("cannot retrieve token:", err.Error())
// 	}
// 	return aws.Credentials{
// 		AccessKeyID:     token,
// 		SecretAccessKey: transport.CellsS3SecretDefault,
// 		SessionToken:    "",
// 	}, nil

// 	// TODO Add cache ?
// 	// appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(token, transport.CellsS3SecretDefault, ""))
// 	// value, err := appCreds.Retrieve(context.TODO())
// 	// if err != nil {
// 	// 	log.Fatal("cannot retrieve credentials:", err.Error())
// 	// }
// 	// return value, nil
// }

// func (s *s3CredentialProvider) IsExpired() bool {
// 	return s.tp.Expired()
// }
