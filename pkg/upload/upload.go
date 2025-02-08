package upload

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/awss3"
	"github.com/redis/go-redis/v9"
)

var pattern = regexp.MustCompile(`https://(.*)\.s3\.(.*)\.amazonaws\.com/(.*)`)

// Redis Prefix Key
const uploadTokenKey = "s3:upload:"

func formatUploadTokenKey(token string) string {
	return uploadTokenKey + token
}

type UploadService struct {
	s3Client    *awss3.Client
	s3Config    *awss3.Config
	redisClient *redis.Client
}

func New(s3Client *awss3.Client, s3Config *awss3.Config, redisClient *redis.Client) *UploadService {
	return &UploadService{
		s3Client:    s3Client,
		s3Config:    s3Config,
		redisClient: redisClient,
	}
}

func (u *UploadService) PresignedUpload(ctx context.Context, stagingPath string, sizeBytes int, fileType string, uploadToken string, tokenData interface{}) (*awss3.PresignedPOST, error) {
	if err := u.setUploadToken(ctx, uploadToken, tokenData); err != nil {
		return nil, fmt.Errorf("can't set upload token: %w", err)
	}

	presignedPost, err := u.s3Client.NewPresignedPost(ctx, &awss3.NewPresignedPostInput{
		Bucket:      u.s3Config.Bucket,
		Key:         stagingPath,
		Size:        sizeBytes,
		ContentType: fileType,
	})
	if err != nil {
		return nil, fmt.Errorf("can't get presigned post: %w", err)
	}

	return presignedPost, nil
}

func (u *UploadService) MoveFileFromStagingToDestination(ctx context.Context, stagingPath, destinationPath string) error {
	err := u.s3Client.CopyToFolder(ctx, u.s3Config.Bucket, stagingPath, destinationPath)
	if err != nil {
		return fmt.Errorf("can't move file from staging to destination: %w", err)
	}

	err = u.s3Client.DeleteObjects(ctx, u.s3Config.Bucket, []string{stagingPath})
	if err != nil {
		return fmt.Errorf("can't delete file from staging: %w", err)
	}
	return nil
}

func (u *UploadService) GetS3Path(destinationPath string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.s3Config.Bucket, u.s3Config.Region, destinationPath)
}

// This function return bucket, region, path, error
func (u *UploadService) SplitS3URLIfValidURL(URL string) (string, string, string, error) {
	matches := pattern.FindStringSubmatch(URL)
	if len(matches) == 4 {
		return matches[1], matches[2], matches[3], nil
	}
	return "", "", "", fmt.Errorf("invalid URL format")
}

func (u *UploadService) DeleteFiles(ctx context.Context, filePath []string) error {
	deleteFiles := make([]string, 0)
	for _, file := range filePath {
		bucket, region, path, err := u.SplitS3URLIfValidURL(file)
		if err == nil && bucket == u.s3Config.Bucket && region == u.s3Config.Region {
			deleteFiles = append(deleteFiles, path)
		}
	}

	if err := u.s3Client.DeleteObjects(ctx, u.s3Config.Bucket, deleteFiles); err != nil {
		return fmt.Errorf("can't delete file from s3: %w", err)
	}

	return nil
}

func (u *UploadService) setUploadToken(ctx context.Context, token string, tokenData interface{}) error {
	marshalData, err := json.Marshal(tokenData)
	if err != nil {
		return fmt.Errorf("can't marshal token data: %w", err)
	}

	key := formatUploadTokenKey(token)
	if err := u.redisClient.Set(ctx, key, marshalData, awss3.EXPIRES_IN).Err(); err != nil {
		return fmt.Errorf("can't set upload token to redis: %w", err)
	}

	return nil
}

func (u *UploadService) GetUploadToken(ctx context.Context, token string) ([]byte, error) {
	key := formatUploadTokenKey(token)
	result, err := u.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("can't get upload token from redis: %w", err)
	}

	return result, nil
}

func (u *UploadService) DeleteUploadToken(ctx context.Context, token string) error {
	key := formatUploadTokenKey(token)
	if err := u.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("can't delete upload token from redis: %w", err)
	}

	return nil
}
