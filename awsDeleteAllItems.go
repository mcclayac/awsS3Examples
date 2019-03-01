package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}


func main() {
	/*if len(os.Args) != 2 {
		exitErrorf("Bucket name required\nUsage: %s BUCKET", os.Args[0])
	}

	bucket := os.Args[1]*/
	bucket := "amcclay-aws-createbucket-sdk"

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	svc := s3.New(sess)

	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
		exitErrorf("Unable to delete objects from bucket %q, %v", bucket, err)
	}

	fmt.Printf("Deleted object(s) from bucket: %s", bucket)


}
