package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pydio/cells-sdk-go/v4/apiv1"
)

func NewCredentialsProvider(sdc *apiv1.SdkConfig, options ...interface{}) (aws.CredentialsProvider, error) {

	var provider aws.CredentialsProvider

	switch sdc.AuthType {
	case apiv1.AuthTypeOAuth:
		provider = &OAuthCredentialsProvider{config: sdc}
	case apiv1.AuthTypePat:
		provider = &PatCredentialsProvider{config: sdc}
	case apiv1.AuthTypeClientAuth:
		provider = &LegacyCredentialsProvider{config: sdc}
	default:
		return nil, fmt.Errorf("unsupported auth type %s, we cannot create a relevant AWS provider", sdc.AuthType)
	}

	// Apply passed options if any of them is relevant here for this class, e.G to set CellsConfigStore.
	for _, o := range options {
		switch typed := o.(type) {
		case apiv1.CredentialProviderOption:
			provider = typed(provider)
		}
	}

	return provider, nil
}
