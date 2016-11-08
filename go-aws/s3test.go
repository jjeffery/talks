
package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.New()
	svc := s3.New(sess)

	input := &s3.ListObjectsInput{ // BlanketyBlankInput, BlanketyBlankOutput
		Bucket: aws.String("bucket-name"), // struct members are pointers
		Prefix: aws.String("prefix/"),
	}

	for {
		output, err := svc.ListObjects(input)
		if err != nil {
			log.Fatal(err)
		}
		for _, obj := range output.Contents {
			fmt.Println(*obj.Key)
		}
		if *output.IsTruncated { // NOTE Pointer
			input.Marker = output.NextMarker
		} else {
			break
		}
	}
}
