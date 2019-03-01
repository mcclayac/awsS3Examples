package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
/*
	if len(os.Args) != 2 {
		exitErrorf("Bucket name missing!\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[1]*/

	bucket := "amcclay-aws-createbucket-sdk"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)


	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully created\n", bucket)

}
