package awsFirstProject

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func main() {

	//sess, err := session.NewSession()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		log.Fatal(err)
	}
	sess = sess
	//fmt.Printf("%+v\n", sess)

	svc := s3.New(sess)
	fmt.Printf("%+v\n", svc)
/*
	// Create Session with MaxRetry configuration to be shared by multiple
	// service clients.
	sess := session.Must(session.NewSession(aws.NewConfig().
		WithMaxRetries(3),
	))

	// Create S3 service client with a specific Region.
	svc := s3.New(sess, aws.NewConfig().
		WithRegion("us-east-2"),
	)
*/





}
