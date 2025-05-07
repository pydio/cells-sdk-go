package s3

import (
	cellssdk "github.com/pydio/cells-sdk-go/v4/apiv1"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/smithy-go/logging"
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

// WithCustomRetry is a helper function that constructs a valid AwsConfigOption
// to fine tune the retry strategy after the transfer of a single part has failed.
//
// maxAttempts is the maximum number of times an operation will be retried before
// giving up. It must be a positive integer, where 1 means only one attempt will
// be made (no retries) and 0 means it will retry indefinitely.
//
// maxBackoffDelay specifies the maximum amount of time to wait between retry
// attempts. At the SDK level, we use an exponential backoff strategy with jitter
// (that introduces randomness in the backoff strategy), so that we wait exponentially
// longer after each failed attempts but no more than the max backoff delay.
//
// extraErrorCodes is a variadic parameter that allows specifying additional
// error codes that should trigger a retry. These are in addition to any error
// codes that the AWS SDK already considers retryable.
func WithCustomRetry(maxAttempts int, maxBackoffDelay time.Duration, extraErrorCodes ...string) cellssdk.AwsConfigOption {
	return func(config aws.Config) aws.Config {
		config.Retryer = func() aws.Retryer {
			tmpR := retry.AddWithMaxAttempts(retry.NewStandard(), maxAttempts)
			if len(extraErrorCodes) > 0 {
				tmpR = retry.AddWithErrorCodes(tmpR, extraErrorCodes...)
			}
			return retry.AddWithMaxBackoffDelay(tmpR, maxBackoffDelay)
		}
		return config
	}
}
