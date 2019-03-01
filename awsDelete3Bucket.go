package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		log.Fatal(err)
	}

	buckets := []string{"amcclay-aws-bucket1", "amcclay-aws-bucket2",
		"amcclay-aws-bucket3", "amcclay-aws-bucket4", "amcclay-aws-bucket5"}
	for _, bucket := range buckets {
		svc := s3.New(sess)
		input := &s3.DeleteBucketInput{
			Bucket: aws.String(bucket),
		}

		result, err := svc.DeleteBucket(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result)
	}
}
