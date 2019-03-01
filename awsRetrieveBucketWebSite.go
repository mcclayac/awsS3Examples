package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
			exitErrorf("bucket name required\nUsage: %s bucket_name", os.Args[0])
		}

		bucket := os.Args[1]

	*/
	bucket := "acloudguru-mcclay-polywebsite"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.GetBucketWebsite(&s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		// Check for the NoSuchWebsiteConfiguration error code telling us
		// that the bucket does not have a website configured.
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchWebsiteConfiguration" {
			exitErrorf("Bucket %s does not have website configuration\n", bucket)
		}
		exitErrorf("Unable to get bucket website config, %v", err)
	}

	fmt.Println("Bucket Website Configuration:")
	fmt.Println(result)
}
