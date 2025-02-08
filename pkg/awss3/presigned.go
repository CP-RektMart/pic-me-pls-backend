package awss3

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cockroachdb/errors"
)

type PresignedPOST struct {
	URL      string            `json:"url"`
	FormData map[string]string `json:"formData"`
}

const XAmzAlgorithm = "AWS4-HMAC-SHA256"

const EXPIRES_IN = 5 * time.Minute // Minutes

type NewPresignedPostInput struct {
	// Key name
	Key string

	Bucket string

	Size int

	ContentType string

	// A list of conditions to include in the policy. Each element can be either a list or a structure.
	// For example:
	// [
	//      {"acl": "public-read"}, ["content-length-range", 2, 5], ["starts-with", "$success_action_redirect", ""]
	// ]
	Conditions []interface{}
}

func (s *Client) NewPresignedPost(ctx context.Context, input *NewPresignedPostInput) (*PresignedPOST, error) {
	// expiration time
	credent, err := s.Presigner.PostPresigner.Cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't get credentials: %w", err)
	}
	expirationTime := time.Now().Add(EXPIRES_IN).UTC()
	dateString := expirationTime.Format("20060102")

	// credentials string
	creds := fmt.Sprintf("%s/%s/%s/s3/aws4_request", credent.AccessKeyID, dateString, s.Presigner.PostPresigner.Region)

	// policy
	policyDoc, err := createPolicyDocument(expirationTime, input.Bucket, input.Key, creds, credent.SessionToken, input.Size, input.ContentType, input.Conditions)
	if err != nil {
		return nil, fmt.Errorf("couldn't create policy document: %w", err)
	}

	// create signature
	signature := createSignature(credent.SecretAccessKey, s.Presigner.PostPresigner.Region, dateString, policyDoc)
	// url
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", input.Bucket, s.Presigner.PostPresigner.Region)

	// expiration time
	dateTimeString := expirationTime.Format("20060102T150405Z")

	// post
	formData := map[string]string{
		"key":              input.Key,
		"Content-Type":     input.ContentType,
		"Policy":           policyDoc,
		"X-Amz-Credential": creds,
		"X-Amz-Signature":  signature,
		"X-Amz-Date":       dateTimeString,
		"X-Amz-Algorithm":  XAmzAlgorithm,
	}
	if credent.SessionToken != "" {
		formData["x-amz-security-token"] = credent.SessionToken
	}
	post := &PresignedPOST{
		URL:      url,
		FormData: formData,
	}

	return post, nil
}

// helpers
func createPolicyDocument(expirationTime time.Time, bucket string, key string, credentialString string, securityToken string, imageSize int, contentType string, extraConditions []interface{}) (string, error) {
	doc := map[string]interface{}{}
	doc["expiration"] = expirationTime.Format("2006-01-02T15:04:05.000Z")

	// conditions
	conditions := []interface{}{}
	conditions = append(conditions, map[string]string{
		"bucket": bucket,
	})
	conditions = append(conditions, []string{
		"eq", "$key", key,
	})
	if securityToken != "" {
		conditions = append(conditions, map[string]string{
			"x-amz-security-token": securityToken,
		})
	}
	conditions = append(conditions, map[string]string{
		"x-amz-credential": credentialString,
	})
	conditions = append(conditions, map[string]string{
		"x-amz-algorithm": XAmzAlgorithm,
	})
	conditions = append(conditions, []interface{}{"content-length-range", imageSize, imageSize})
	conditions = append(conditions, []interface{}{"eq", "$Content-Type", contentType})
	conditions = append(conditions, map[string]string{
		"x-amz-date": expirationTime.Format("20060102T150405Z"),
	})

	// other conditions
	conditions = append(conditions, extraConditions...)

	doc["conditions"] = conditions
	// base64 encoded json string
	jsonBytes, err := json.Marshal(doc)
	if err != nil {
		return "", fmt.Errorf("couldn't marshal policy document: %w", err)
	}

	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

func createSignature(secretKey string, region string, dateString string, stringToSign string) string {
	// Helper to make the HMAC-SHA256.
	makeHmac := func(key []byte, data []byte) []byte {
		hash := hmac.New(sha256.New, key)
		hash.Write(data)
		return hash.Sum(nil)
	}

	h1 := makeHmac([]byte("AWS4"+secretKey), []byte(dateString))
	h2 := makeHmac(h1, []byte(region))
	h3 := makeHmac(h2, []byte("s3"))
	h4 := makeHmac(h3, []byte("aws4_request"))
	signature := makeHmac(h4, []byte(stringToSign))
	return hex.EncodeToString(signature)
}

// GetObject makes a presigned request that can be used to get an object from a bucket.
// The presigned request is valid for the specified number of seconds.
func (s *Client) GetObject(ctx context.Context,
	bucketName string, objectKey string,
) (*v4.PresignedHTTPRequest, error) {
	request, err := s.Presigner.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	},
	)
	if err != nil {
		return nil, errors.Newf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, nil
}

// DeleteObject makes a presigned request that can be used to delete an object from a bucket.
func (s *Client) DeleteObject(ctx context.Context, bucketName string, objectKey string) (*v4.PresignedHTTPRequest, error) {
	request, err := s.Presigner.PresignClient.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, errors.Newf("Couldn't get a presigned request to delete object %v. Here's why: %v\n", objectKey, err)
	}

	return request, fmt.Errorf(fmt.Sprintf("couldn't get a presigned request to delete %v", objectKey), err)
}
