package awss3

import (
	"context"
	"fmt"
	"net/url"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cockroachdb/errors"
)

type Config struct {
	Bucket          string `env:"BUCKET"`
	Region          string `env:"REGION"`
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
}

type Client struct {
	Client    *s3.Client
	Config    *Config
	Presigner *Presigner
}

type Presigner struct {
	PresignClient *s3.PresignClient
	PostPresigner *PostPresigner
}

type PostPresigner struct {
	Region string
	Cfg    aws.Config
}

func NewS3(ctx context.Context, config *Config) (*Client, error) {
	cfg := getS3Config(ctx, config)
	s3Client := s3.NewFromConfig(cfg)
	Presigner := &Presigner{
		s3.NewPresignClient(s3Client),
		&PostPresigner{
			Region: config.Region,
			Cfg:    cfg,
		},
	}

	return &Client{
		Client:    s3Client,
		Config:    config,
		Presigner: Presigner,
	}, nil
}

func getS3Config(ctx context.Context, config *Config) aws.Config {
	var cfg aws.Config
	var err error
	if config.AccessKeyID != "" && config.SecretAccessKey != "" {
		cfg, err = awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(config.Region),
			awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.AccessKeyID, config.SecretAccessKey, "")),
		)
		if err != nil {
			logger.Panic(err.Error())
		}
		return cfg
	}

	appCreds := aws.NewCredentialsCache(ec2rolecreds.New())
	_, err = appCreds.Retrieve(ctx)
	if err == nil {
		cfg, err = awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(config.Region), awsConfig.WithCredentialsProvider(appCreds))
	} else {
		cfg, _ = awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(config.Region))
		_, err = cfg.Credentials.Retrieve(ctx)
	}

	if err != nil {
		logger.Panic(err.Error())
	}
	return cfg
}

func (s *Client) CopyToFolder(ctx context.Context, bucketName string, filePath string, targetPath string) error {
	escapedFilePath := url.PathEscape(filePath)

	_, err := s.Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(bucketName),
		CopySource: aws.String(fmt.Sprintf("%v/%v", bucketName, escapedFilePath)),
		Key:        aws.String(fmt.Sprintf("%v", targetPath)),
	})
	if err != nil {
		return errors.Newf("Couldn't copy object from %v:%v to %v:%v. Here's why: %v\n",
			bucketName, filePath, bucketName, targetPath, err)
	}
	return errors.Wrap(err, "Couldn't copy object")
}

func (s *Client) ListObjectsID(ctx context.Context, bucketName string, folderPath string) ([]string, error) {
	result, err := s.Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(folderPath),
	})
	var contents []types.Object
	if err != nil {
		return nil, errors.Newf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		contents = result.Contents
	}
	keys := make([]string, 0, len(contents))
	for _, content := range contents {
		keys = append(keys, *content.Key)
	}
	return keys, errors.Wrap(err, "Couldn't list objects")
}

func (s *Client) DeleteObjects(ctx context.Context, bucketName string, objectKeys []string) error {
	objectIds := make([]types.ObjectIdentifier, 0)
	for _, key := range objectKeys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}
	_, err := s.Client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		return errors.Newf("Couldn't delete objects from bucket %v. Here's why: %v\n", bucketName, err)
	}
	return errors.Wrap(err, "Couldn't delete objects")
}
