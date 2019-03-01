package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

func main() {




	svc := s3.New(session.New())
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("twistlock.jpg")),
		//Body:   aws.ReadSeekCloser(strings.NewReader("/Users/mcclayac/Google Drive/images/BlackSalesGirl.jpg")),
		Bucket: aws.String("amcclay-awsgobucket"),
		Key:    aws.String("twistlock.jpg"),
	}

	result, err := svc.PutObject(input)
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
