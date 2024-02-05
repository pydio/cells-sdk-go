package s3

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/logging"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
)

// WithCellsConfigStore is the entry point to provide an external store that exposes
// a method to refresh and store credentials and become then the unique source of truth
// to retrieve current credentials.
func WithCellsConfigStore(store cellssdk.ConfigRefresher) cellssdk.CredentialProviderOption {
	return func(provider aws.CredentialsProvider) aws.CredentialsProvider {
		if cs, ok := provider.(cellssdk.ConfigRefresherConsumer); ok {
			cs.SetConfigRefresher(store)
		}
		return provider
	}
}

// WithLogger is a helper function to construct a valid AwsConfigOption
// to define and configure the logging strategy that will be used by the AWS SDK
// when performing file transfers.
func WithLogger(writer io.Writer, logMode aws.ClientLogMode) cellssdk.AwsConfigOption {
	return func(config aws.Config) aws.Config {
		config.Logger = logging.NewStandardLogger(writer)
		config.ClientLogMode = logMode
		return config
	}
}

// WithRegion is a helper function to define a specific Region when talking
// with S3 remote server. This is usually useless when remote server is Cells.
func WithRegion(region string) cellssdk.AwsConfigOption {
	return func(config aws.Config) aws.Config {
		config.Region = region
		return config
	}
}
