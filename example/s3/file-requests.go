package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pydio/cells-sdk-go/config"
)

func GetFile(workspaceSlug, pathToFile string) (*s3.GetObjectOutput, error) {
	svc := s3.New(config.DefaultS3Session)
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
	svc := s3.New(config.DefaultS3Session)
	obj, err := svc.PutObject((&s3.PutObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile).
		SetBody(content),
	)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func DeleteFile(workspaceSlug, pathToFile string) (*s3.DeleteObjectOutput, error) {
	svc := s3.New(config.DefaultS3Session)

	obj, err := svc.DeleteObject((&s3.DeleteObjectInput{}).
		SetBucket(config.DefaultS3Config.Bucket).
		SetKey(workspaceSlug + "/" + pathToFile),
	)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
