package s3

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pydio/cells-sdk-go/config"
)

func GetFile(workspaceSlug, pathToFile string) (*s3.GetObjectOutput, error) {

	svc, err := config.GetPreparedS3CLient(config.DefaultConfig, config.DefaultS3Config)
	if err != nil {
		return nil, err
	}

	obj, err := svc.GetObject((&s3.GetObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile),
	)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func PutFile(workspaceSlug, pathToFile string, content io.ReadSeeker) (*s3.PutObjectOutput, error) {

	svc, err := config.GetPreparedS3CLient(config.DefaultConfig, config.DefaultS3Config)
	if err != nil {
		return nil, err
	}

	obj, err := svc.PutObject((&s3.PutObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile).
		SetBody(content),
	)
	if err != nil {
		fmt.Println("Could not put file: " + workspaceSlug + "/" + pathToFile)
		return nil, err
	}
	return obj, nil
}

func DeleteFile(workspaceSlug, pathToFile string) (*s3.DeleteObjectOutput, error) {

	svc, err := config.GetPreparedS3CLient(config.DefaultConfig, config.DefaultS3Config)
	if err != nil {
		return nil, err
	}

	obj, err := svc.DeleteObject((&s3.DeleteObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile),
	)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
