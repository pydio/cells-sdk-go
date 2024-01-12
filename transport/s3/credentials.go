package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pydio/cells-sdk-go/v4/transport"
)

type s3CredentialProvider struct {
	tp transport.TokenProvider
}

func (s *s3CredentialProvider) Retrieve(_ context.Context) (aws.Credentials, error) {

	token, err := s.tp.Retrieve()
	if err != nil {
		log.Fatal("cannot retrieve token:", err.Error())
	}
	return aws.Credentials{
		AccessKeyID:     token,
		SecretAccessKey: transport.CellsS3SecretDefault,
		SessionToken:    "",
	}, nil

	// TODO Add cache ?
	// appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(token, transport.CellsS3SecretDefault, ""))
	// value, err := appCreds.Retrieve(context.TODO())
	// if err != nil {
	// 	log.Fatal("cannot retrieve credentials:", err.Error())
	// }
	// return value, nil
}

func (s *s3CredentialProvider) IsExpired() bool {
	return s.tp.Expired()
}

func AsS3CredentialsProvider(provider transport.TokenProvider) aws.CredentialsProvider {
	return &s3CredentialProvider{
		tp: provider,
	}
}
