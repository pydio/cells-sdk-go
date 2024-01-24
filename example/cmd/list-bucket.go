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

var listBuckets = &cobra.Command{
	Use: "list-buckets",
	Long: `Simply list all accessible buckets
	
	Example:
	$ go run main.go list-buckets -u https://files.example.com -l admin -p your-password 
`,
	Run: func(cmd *cobra.Command, args []string) {
		s3Conf := getS3ConfigFromSdkConfig(DefaultConfig)
		client, e := ts3.GetClient(newDummyStore(), DefaultConfig, s3Conf)
		if e != nil {
			log.Fatal(e)
		}
		// Adapt below and rather use this to test on local client.
		// client = localMinioClient()
		doTest(cmd, client)

	},
}

func doTest(cmd *cobra.Command, client *s3.Client) {
	o, e := client.ListBuckets(
		context.TODO(),
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

// localMinioClient simply perform a comparative test on a local minio server
func localMinioClient() *s3.Client {

	localUrl := "http://192.168.0.136:9000"
	localAccessKey := "nhlXm2iabssGndbuJ1ei"
	localSecret := "WCdpoRdBTv6UPrNvRGYUa5oTmBjh2rg6Er2APpTb"

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(&MyProvider{
			AccessKey: localAccessKey,
			Secret:    localSecret,
		}),
	)
	if err != nil {
		log.Fatal("cannot load default S3 session configuration:", err.Error())
	}

	fo := func(o *s3.Options) {
		o.BaseEndpoint = aws.String(localUrl)
	}

	return s3.NewFromConfig(cfg, fo)
}

type MyProvider struct {
	AccessKey string
	Secret    string
}

func (m *MyProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
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
