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
	/*if len(os.Args) != 2 {
		exitErrorf("bucket name required\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[1]*/
	bucket := "amcclay-aws-createbucket-sdk"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait until bucket is deleted before finishing
	fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully deleted\n", bucket)

}
