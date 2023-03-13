package s3

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pydio/cells-sdk-go/v3/transport"
)

type s3CredentialProvider struct {
	tp transport.TokenProvider
}

func (s *s3CredentialProvider) Retrieve() (credentials.Value, error) {
	token, err := s.tp.Retrieve()
	return credentials.Value{
		AccessKeyID:     token,
		SecretAccessKey: transport.CellsS3SecretDefault,
	}, err
}

func (s *s3CredentialProvider) IsExpired() bool {
	return s.tp.Expired()
}

func AsS3CredentialsProvider(provider transport.TokenProvider) credentials.Provider {
	return &s3CredentialProvider{
		tp: provider,
	}
}
