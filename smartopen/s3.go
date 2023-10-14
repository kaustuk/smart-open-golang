package smartopen

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var smartOpenS3Obj smartOpenS3Client

type smartOpenS3Client struct {
	s3Client *s3.Client
}

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	s3Client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://0.0.0.0:4566")
	})

	smartOpenS3Obj.s3Client = s3Client
}

func (s smartOpenS3Client) openFile(bucket string, objectKey string) (io.ReadCloser, error) {
	fmt.Printf("%s \t %s\n", bucket, objectKey)
	result, err := s.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}
