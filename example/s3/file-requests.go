package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pydio/cells-sdk-go/config"
)

func GetFile(workspaceSlug, pathToFile string, headOnly bool) error {

	svc := s3.New(config.DefaultS3Session)

	obj, err := svc.GetObject((&s3.GetObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile),
	)
	if err != nil {
		return err
	}

	fmt.Println("Got an object " + obj.String())
	return nil
}
