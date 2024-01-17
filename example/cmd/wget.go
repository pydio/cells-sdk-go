package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/transport"
	ts3 "github.com/pydio/cells-sdk-go/v4/transport/s3"
	"github.com/spf13/cobra"
)

var (
	getPath string
)

type dummyConfigStore struct{}

func (pcp *dummyConfigStore) RefreshIfRequired(*cells_sdk.SdkConfig) (bool, error) {
	return false, nil
}

func newDummyStore() transport.ConfigStore {
	return &dummyConfigStore{}
}

var getFile = &cobra.Command{
	Use: "wget",
	Long: `Get a file from your Pydio Cells server.
	
	Example:
	$ go run main.go wget -u https://pydio.example.com -l admin -p password -f common-files/image.jpg  
`,
	Run: func(cmd *cobra.Command, args []string) {

		s3Conf := getS3ConfigFromSdkConfig(*DefaultConfig)
		client, e := ts3.GetClient(newDummyStore(), DefaultConfig, &s3Conf)
		if e != nil {
			log.Fatal(e)
		}
		output, e := client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: &s3Conf.Bucket,
			Key:    &getPath,
		})
		if e != nil {
			log.Fatal(e)
		}
		cmd.Println("Downloading file from server...")
		base := path.Base(getPath)
		target, e := os.OpenFile(filepath.Join(".", base), os.O_CREATE|os.O_WRONLY, 0755)
		if e != nil {
			log.Fatal(e)
		}
		defer target.Close()
		written, e := io.Copy(target, output.Body)
		if e != nil {
			log.Fatal(e)
		} else {
			cmd.Printf("Written %d bytes to file %s\n", written, base)
		}
	},
}

func init() {
	getFile.Flags().StringVarP(&getPath, "filepath", "f", "", "File path")
	ExampleCmd.AddCommand(getFile)
}
