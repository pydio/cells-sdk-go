package cmd

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"

	ts3 "github.com/pydio/cells-sdk-go/v5/transport/s3"
)

const (
	// Adapt the below values to test locally, e.G with a local minio server.
	testLocally    = false
	localUrl       = "http://192.168.0.10:9000"
	localRegion    = "us-east-1"
	localAccessKey = "nhlXm3iabssGndbuJ1ei"
	localSecret    = "WCdpoRppTv6UPrNvRGYUa5oTmBjh2rg6Er2APpTb"
)

var listBuckets = &cobra.Command{
	Use: "list-buckets",
	Long: `Simply list all accessible buckets
	
	Example:
	$ go run main.go list-buckets -u https://files.example.com -l admin -p your-password 
`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, e := ts3.LoadConfig(cmd.Context(), DefaultConfig)
		if e != nil {
			log.Fatal(e)
		}
		client := ts3.NewClientFromConfig(cfg, DefaultConfig.Url)

		if testLocally {
			client = localMinioClient(cmd.Context())
		}
		doTest(cmd, client)
	},
}

func doTest(cmd *cobra.Command, client *s3.Client) {
	o, e := client.ListBuckets(
		cmd.Context(),
		&s3.ListBucketsInput{},
	)
	if e != nil {
		log.Fatal(e)
	}
	cmd.Printf("Found %d buckets:\n", len(o.Buckets))
	for _, b := range o.Buckets {
		cmd.Println("\t-", *b.Name, "- created at:", b.CreationDate)
	}
}

// localMinioClient simply perform a comparative test on a local minio server.
// Adapt the hard-coded constants (see above) to fit with your test setup.
func localMinioClient(ctx context.Context) *s3.Client {

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(localRegion),
		config.WithCredentialsProvider(&MyProvider{
			AccessKey: localAccessKey,
			Secret:    localSecret,
		}),
	)
	if err != nil {
		log.Fatal("cannot load default S3 session configuration:", err.Error())
	}

	return ts3.NewClientFromConfig(cfg, localUrl)

	//fo := func(o *s3.Options) {
	//	o.BaseEndpoint = aws.String(localUrl)
	//}
	//
	//return s3.NewFromConfig(cfg, fo)
}

type MyProvider struct {
	AccessKey string
	Secret    string
}

func (m *MyProvider) Retrieve(_ context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     m.AccessKey,
		SecretAccessKey: m.Secret,
	}, nil
}

func (m *MyProvider) IsExpired() bool {
	return false
}

func init() {
	ExampleCmd.AddCommand(listBuckets)
}
