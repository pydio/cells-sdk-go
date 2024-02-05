package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

func NewCredentialsProvider(sdc *cellssdk.SdkConfig, options ...interface{}) (aws.CredentialsProvider, error) {

	var provider aws.CredentialsProvider

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
