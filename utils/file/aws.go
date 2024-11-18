package aws

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"orchid.admin.service/conf"
	"orchid.admin.service/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"golang.org/x/exp/slog"
)

type AwsBucket struct {
	c      *conf.Config
	client *s3.Client
}

func NewAwsBucket(c *conf.Config) (*AwsBucket, error) {
	client := s3.New(s3.Options{
		Region:      c.AwsS3.Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(c.AwsS3.Key, c.AwsS3.Secret, "")),
		EndpointResolver: s3.EndpointResolverFromURL(
			fmt.Sprintf("https://s3.%s.amazonaws.com", c.AwsS3.Region),
		),
		UsePathStyle: true, // Ensure path-style URLs for compatibility
	})

	aws3 := &AwsBucket{c, client}
	isExist, err := aws3.BucketExists()
	if err != nil {
		slog.Error("Unable to check if bucket exists", slog.Any("err", err))
		return nil, err
	}

	if !isExist {
		err = aws3.CreateBucket(utils.BucketName)
		if err != nil {
			slog.Error("Unable to create bucket", slog.Any("err", err))
			return nil, err
		}
	}

	return aws3, nil
}

func (awsb *AwsBucket) BucketExists() (bool, error) {
	exists := true

	_, err := awsb.client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(utils.BucketName),
	})
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				exists = false
				err = nil
			default:
				slog.Error("error on bucket", slog.Any("err", err))
			}
		}
	}

	return exists, err
}

func (awsb *AwsBucket) CreateBucket(name string) error {
	_, err := awsb.client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint("ap-southeast-1"),
		},
	})
	if err != nil {
		slog.Error("Couldn't create bucket", slog.Any("err", err))
		return err
	}

	return nil
}

func (awsb *AwsBucket) UploadFile(objectKey string, file io.Reader) error {
	_, err := awsb.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(utils.BucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		slog.Error("Couldn't upload file", slog.Any("err", err))
		return err
	}

	return nil
}

func (awsb *AwsBucket) DownloadFile(objectKey string) ([]byte, error) {
	result, err := awsb.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(utils.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		slog.Error("Couldn't get object", slog.Any("err", err))
		return nil, err
	}

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		slog.Error("Couldn't read object ", slog.Any("err", err))
		return nil, err
	}

	return body, nil
}

func (awsb *AwsBucket) DeleteObject(objectKey string) error {
	var objectIds []types.ObjectIdentifier
	objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(objectKey)})
	res, err := awsb.client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(utils.BucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if len(res.Errors) > 0 {
		slog.Error("res err", slog.Any("Err", err))
		return err
	}
	if err != nil {
		slog.Error("Couldn't delete object from bucket", slog.Any("objectKey", objectKey), slog.Any("err", err))
		return err
	}

	return nil
}

func (awsb *AwsBucket) GetPresignedObject(contentDisp, file *string, objectKey string, lifetimeSecs int64) (string, error) {
	presignClient := s3.NewPresignClient(awsb.client)
	request, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket:                     aws.String(utils.BucketName),
		Key:                        aws.String(objectKey),
		ResponseContentDisposition: contentDisp,
		ResponseContentType:        file,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs) * time.Second
	})
	if err != nil {
		slog.Error("Couldn't get a presigned request to get", slog.Any("err", err))
		return "", err
	}

	return request.URL, nil
}
