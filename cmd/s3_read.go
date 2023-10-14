package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	fmt.Printf("hello from main\n")

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	s3Client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://0.0.0.0:4566")
	})

	bucketName := "test-kaustuk-bucket"
	objectKey := "testfile"
	fileName := "kaustuk-download"

	result, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return
	}
	defer result.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	// return err

	// s3Client := s3.NewFromConfig(sdkConfig)

	// result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// if err != nil {
	// 	fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	// 	return
	// }
	// if len(result.Buckets) == 0 {
	// 	fmt.Println("You don't have any buckets!")
	// } else {

	// 	for _, bucket := range result.Buckets {
	// 		fmt.Printf("\t%v\n", *bucket.Name)
	// 	}
	// }

}
