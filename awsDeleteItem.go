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
	/*if len(os.Args) != 3 {
		exitErrorf("Bucket and object name required\nUsage: %s bucket_name object_name",
			os.Args[0])
	}

	bucket := os.Args[1]
	obj := os.Args[2]*/


	bucket := "amcclay-awscreatebucket-sdk"
	obj := "BlackSalesGirl.jpg"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})
	if err != nil {
		exitErrorf("Unable to delete object %q from bucket %q, %v", obj, bucket, err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	fmt.Printf("Object %q successfully deleted\n", obj)

}
