package cmd

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/v4/transport/s3"
)

var listBuckets = &cobra.Command{
	Use: "list-buckets",
	Long: `Simply list root buckets for passed user.
	
	Example:
	$ go run main.go list-buckets -u http://192.168.0.136:9000 -l admin -p your-password 
`,
	Run: func(cmd *cobra.Command, args []string) {

		s3Conf := getS3ConfigFromSdkConfig(*DefaultConfig)
		client, e := s3.GetClient(DefaultConfig, &s3Conf)
		if e != nil {
			log.Fatal(e)
		}

		// Adapt below and rather use this to test on local client.
		// client := localMinioClient()
		doTest(cmd, client)
	},
}

func doTest(cmd *cobra.Command, client *aws_s3.S3) {
	o, e := client.ListBuckets(
		&aws_s3.ListBucketsInput{},
	)
	if e != nil {
		log.Fatal(e)
	}
	cmd.Printf("Found %d buckets:\n", len(o.Buckets))
	for _, b := range o.Buckets {
		cmd.Println("\t -", *b.Name, "-", b.CreationDate)
	}
}

// localMinioClient simply perform a comparative test on a local minio server
func localMinioClient() *aws_s3.S3 {

	localUrl := "http://192.168.0.136:9000"
	localAccessKey := "nhlXm2iabssGndbuJ1ei"
	localSecret := "WCdpoRdBTv6UPrNvRGYUa5oTmBjh2rg6Er2APpTb"

	conf := aws.NewConfig()
	conf.WithRegion("us-east-1")
	conf.WithEndpoint(localUrl)
	conf.WithCredentials(credentials.NewCredentials(&MyProvider{
		AccessKey: localAccessKey,
		Secret:    localSecret,
	}))
	s3Session, err := session.NewSession(conf)
	if err != nil {
		log.Fatal("cannot initialise default S3 session: " + err.Error())
	}
	return aws_s3.New(s3Session)
}

type MyProvider struct {
	AccessKey string
	Secret    string
}

func (m *MyProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
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
