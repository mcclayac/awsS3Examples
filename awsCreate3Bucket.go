package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	svc := s3.New(session.New())



	buckets := []string{"amcclay-aws-bucket1", "amcclay-aws-bucket2",
		"amcclay-aws-bucket3", "amcclay-aws-bucket4", "amcclay-aws-bucket5"}
	for _, bucket := range buckets {
		input := &s3.CreateBucketInput{
			Bucket: aws.String(bucket),
			CreateBucketConfiguration: &s3.CreateBucketConfiguration{
				LocationConstraint: aws.String("us-east-2"),
			},
		}

		result, err := svc.CreateBucket(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeBucketAlreadyExists:
					fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
				case s3.ErrCodeBucketAlreadyOwnedByYou:
					fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
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
