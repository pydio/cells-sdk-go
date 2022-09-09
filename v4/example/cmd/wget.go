package cmd

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	aws "github.com/aws/aws-sdk-go/service/s3"
	"github.com/pydio/cells-sdk-go/v4/transport/s3"
	"github.com/spf13/cobra"
)

var (
	getPath string
)

var getFile = &cobra.Command{
	Use: "wget",
	Long: `Get a file from your Pydio Cells server.
	
	Example:
	$ go run main.go -u https://pydio.example.com -l admin -p password -f common-files/image.jpg wget 
`,
	Run: func(cmd *cobra.Command, args []string) {
		s3Conf := getS3ConfigFromSdkConfig(*DefaultConfig)
		cl, e := s3.GetClient(DefaultConfig, &s3Conf)
		if e != nil {
			log.Fatal(e)
		}
		output, e := cl.GetObject(&aws.GetObjectInput{
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
