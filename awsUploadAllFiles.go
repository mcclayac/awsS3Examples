package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"log"
	"os"
)


func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {


	directoryRoot := "/Users/mcclayac/Google Drive/images"
	bucket := "amcclay-aws-createbucket-sdk"

	files, err := ioutil.ReadDir(directoryRoot)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Mode(), file.Name())

		if file.IsDir() {
			//fmt.Println(file.Name() + " is Directory")
			continue
		}

		//filename := os.Args[2]
		//filename := "bigFoot.jpg"
		keyName := file.Name()
		filename := directoryRoot + "/" + keyName


		file, err := os.Open(filename)
		if err != nil {
			exitErrorf("Unable to open file %q, %v", err)
		}

		//defer file.Close()

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-2")},
		)

		// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
		// for more information on configuring part size, and concurrency.
		//
		// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
		uploader := s3manager.NewUploader(sess)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key: aws.String(keyName),
			Body: file,
		})
		if err != nil {
			// Print the error and exit.
			exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		}

		fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
		file.Close()
	}
}
