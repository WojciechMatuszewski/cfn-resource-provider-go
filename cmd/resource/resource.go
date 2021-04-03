package resource

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	currentModel.ResourceId = aws.String("wojtek-super-bucket")

	log.Println("Creating a bucket with resourceId", *currentModel.ResourceId)

	if currentModel.BucketName == nil {
		err := errors.New("BucketName is a required property")
		return handler.NewFailedEvent(err), err
	}

	bucketName := getBucketName(*currentModel.BucketName)
	log.Println("BucketName of the currentModel will be", bucketName)
	currentModel.BucketName = aws.String(bucketName)

	svc := s3.New(req.Session)
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		response := handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
			ResourceModel:   currentModel,
		}

		return response, nil
	}

	response := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}

	return response, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("Not implemented: Read")
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("Not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.BucketName == nil {
		log.Println("Integrity issue, BucketName is nil")

		err := errors.New("BucketName is required")
		return handler.NewFailedEvent(err), err
	}

	bucketName := getBucketName(*currentModel.BucketName)
	log.Println("Deleting bucket with a name of", bucketName)

	svc := s3.New(req.Session)
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})

	if err != nil {
		return handler.NewFailedEvent(err), err
	}

	currentModel.BucketName = nil
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Bucket deleted",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

/*
	The `currentModel` points to the non-transformed values.
	The `bucketName` is the bucket name of that the user provided
*/
func getBucketName(providedName string) string {
	normalized := strings.ToLower(providedName)

	bucketName := normalized
	if len(normalized) >= 27 {
		bucketName = normalized[0:27]
	}

	return fmt.Sprintf("%s-%x", bucketName, md5.Sum([]byte(bucketName)))
}
