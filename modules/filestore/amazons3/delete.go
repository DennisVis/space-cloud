package amazons3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// DeleteFile deletes a file from S3
func (a *AmazonS3) DeleteFile(ctx context.Context, project, path string) error {
	svc := s3.New(a.client)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(project), Key: aws.String(path)})
	if err != nil {
		return err
	}

	return svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(project),
		Key:    aws.String(project + path),
	})
}

// DeleteDir deletes a directory in S3
func (a *AmazonS3) DeleteDir(ctx context.Context, project, path string) error {
	// TODO: Consider AWS operation limit
	svc := s3.New(a.client)

	// Setup BatchDeleteIterator to iterate through a list of objects.
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(project),
		Prefix: aws.String(path),
	})

	// Traverse iterator deleting each object
	return s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
}
