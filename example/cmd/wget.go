package cmd

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	ts3 "github.com/pydio/cells-sdk-go/v5/transport/s3"
)

var (
	getPath string
)

type dummyConfigStore struct{}

func (pcp *dummyConfigStore) RefreshIfRequired(*cells_sdk.SdkConfig) (bool, error) {
	return false, nil
}

func newDummyStore() cells_sdk.ConfigStore {
	return &dummyConfigStore{}
}

var getFile = &cobra.Command{
	Use: "wget",
	Long: `Get a file from your Pydio Cells server.
	
	Example:
	$ go run main.go wget -u https://pydio.example.com -l admin -p password -f common-files/image.jpg  
`,
	Run: func(cmd *cobra.Command, args []string) {

		cfg, e := ts3.LoadConfig(cmd.Context(), DefaultConfig)
		if e != nil {
			log.Fatal(e)
		}
		client := ts3.NewClientFromConfig(cfg, DefaultConfig.Url)

		output, e := client.GetObject(cmd.Context(), &s3.GetObjectInput{
			Bucket: aws.String(cells_sdk.DefaultS3Bucket),
			Key:    aws.String(getPath),
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
