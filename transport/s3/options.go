package s3

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/logging"

	"github.com/pydio/cells-sdk-go/v5"
)

// WithCellsConfigStore is the entry point to provide an external store that exposes
// a method to refresh and store credentials and become then the unique source of truth
// to retrieve current credentials.
func WithCellsConfigStore(store cells_sdk.ConfigStore) cells_sdk.CredentialProviderOption {
	return func(provider cells_sdk.CellsCredentialsProvider) cells_sdk.CellsCredentialsProvider {
		provider.SetConfigStore(store)
		return provider
	}
}

// WithLogger is a helper function to construct a valid AwsConfigOption
// to define and configure the logging strategy that will be used by the AWS SDK
// when performing file transfers.
func WithLogger(writer io.Writer, logMode aws.ClientLogMode) cells_sdk.AwsConfigOption {
	return func(config aws.Config) aws.Config {
		config.Logger = logging.NewStandardLogger(writer)
		config.ClientLogMode = logMode
		return config
	}
}

// WithRegion is a helper function to define a specific Region when talking
// with S3 remote server. This is usually useless when remote server is Cells.
func WithRegion(region string) cells_sdk.AwsConfigOption {
	return func(config aws.Config) aws.Config {
		config.Region = region
		return config
	}
}
