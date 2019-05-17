package transport

import (
	"context"
	"io"
	"net/url"

	sdk "github.com/pydio/cells-sdk-go"
	"github.com/pydio/cells/common/proto/tree"
	"github.com/pydio/cells/common/views"
	"github.com/pydio/minio-go"
)

type S3Client struct {
	config   *sdk.SdkConfig
	s3config *sdk.S3Config
}

func NewS3Client(config *sdk.SdkConfig) *S3Client {
	return &S3Client{
		config: config,
		s3config: &sdk.S3Config{
			Bucket:                 "data",
			ApiKey:                 "gateway",
			ApiSecret:              "gatewaysecret",
			UsePydioSpecificHeader: false,
			IsDebug:                false,
			Region:                 "us-east-1",
			Endpoint:               config.Url,
		},
	}
}

func (g *S3Client) GetObject(ctx context.Context, node *tree.Node, requestData *views.GetRequestData) (io.ReadCloser, error) {
	jwt, err := RetrieveToken(g.config)
	if err != nil {
		return nil, err
	}
	u, _ := url.Parse(g.s3config.Endpoint)
	mc, e := minio.NewCore(u.Host, jwt, g.s3config.ApiSecret, u.Scheme == "https")
	if e != nil {
		return nil, e
	}
	r, _, e := mc.GetObject(g.s3config.Bucket, node.Path, minio.GetObjectOptions{})
	return r, e
}

func (g *S3Client) PutObject(ctx context.Context, node *tree.Node, reader io.Reader, requestData *views.PutRequestData) (int64, error) {

	jwt, err := RetrieveToken(g.config)
	if err != nil {
		return 0, err
	}
	u, _ := url.Parse(g.s3config.Endpoint)
	mc, e := minio.NewCore(u.Host, jwt, g.s3config.ApiSecret, u.Scheme == "https")
	if e != nil {
		return 0, e
	}
	return mc.PutObjectWithContext(ctx, g.s3config.Bucket, node.Path, reader, requestData.Size, minio.PutObjectOptions{
		UserMetadata: requestData.Metadata,
	})
}

func (g *S3Client) CopyObject(ctx context.Context, from *tree.Node, to *tree.Node, requestData *views.CopyRequestData) (int64, error) {
	jwt, err := RetrieveToken(g.config)
	if err != nil {
		return 0, err
	}
	mc, e := minio.New(g.s3config.Endpoint, g.s3config.ApiKey, jwt, false)
	if e != nil {
		return 0, e
	}
	dst, e := minio.NewDestinationInfo(g.s3config.Bucket, to.Path, nil, requestData.Metadata)
	if e != nil {
		return 0, e
	}
	src := minio.NewSourceInfo(g.s3config.Bucket, from.Path, nil)
	return 0, mc.CopyObject(dst, src)
}
