package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Controller struct {
	S3 *s3.S3
}

// createSessionForLocalstack Create a session for localstack
func createSessionForLocalstack(region string, endpoint string, profile string) *session.Session {
	conf := aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess, _ := session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		Config:            conf,
		SharedConfigState: session.SharedConfigEnable,
	})
	return sess
}

func CreateS3Controller(region string, endpoint string, profile string) S3Controller {
	sess := createSessionForLocalstack(
		region,
		endpoint,
		profile,
	)

	s3Ctr := S3Controller{
		S3: s3.New(sess),
	}

	return s3Ctr
}

func (s3Ctr *S3Controller) CreateBuckets(buckets []string) error {
	for _, b := range buckets {
		_, err := s3Ctr.S3.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(b),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s3Ctr *S3Controller) ListBuckets() error {
	result, err := s3Ctr.S3.ListBuckets(nil)
	if err != nil {
		return err
	}
	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return nil
}

func main() {
	fmt.Printf("")
	region := "ap-northeast-1"
	endpoint := "http://localstack:4572"
	profile := "localstack"
	buckets := []string{
		"sample1",
		"sample2",
		"sample3",
	}

	s3Ctr := CreateS3Controller(region, endpoint, profile)

	s3Ctr.CreateBuckets(buckets)
	s3Ctr.ListBuckets()

}
