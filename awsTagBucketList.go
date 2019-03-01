package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Tag S3 bucket MyBucket with cost center tag "123456" and stack tag "MyTestStack".
//
// See:
//    http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func main() {
	// Pre-defined values
	bucket := "amcclay-awsgobucket"

	/*tagName1 := "Cost Center"
	tagValue1 := "123456"
	tagName2 := "Stack"
	tagValue2 := "MyTestStack"
*/
	// Initialize a session in us-west-2 that the SDK will use to load credentials
	// from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create S3 service client
	svc := s3.New(sess)


	// Now show the tags
	// Create input for GetBucket method
	input := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketTagging(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numTags := len(result.TagSet)

	if numTags > 0 {
		fmt.Println("Found", numTags, "Tag(s):")
		fmt.Println("")

		for _, t := range result.TagSet {
			fmt.Println("  Key:  ", *t.Key)
			fmt.Println("  Value:", *t.Value)
			fmt.Println("")
		}
	} else {
		fmt.Println("Did not find any tags")
	}
}

