package uploadprovider

import (
	"bytes"
	"context"
	"flag"
	"fmt"

	"social-todo-list/common"

	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type s3Provider struct {
	name       string
	bucketName string
	region     string
	accessKey  string
	secretKey  string
	endPoint   string
	domain     string
	client     *s3.Client
	logger     logger.Logger
}

func NewS3Provider(name string) *s3Provider {
	return &s3Provider{
		name: name,
	}
}

func (provider *s3Provider) GetPrefix() string {
	return provider.Name()
}

func (provider *s3Provider) Get() interface{} {
	return provider
}

func (provider *s3Provider) Name() string {
	return provider.name
}

func (provider *s3Provider) InitFlags() {
	flag.StringVar(&provider.accessKey, "storage-access-key", "", "Cloud storage access key")
	flag.StringVar(&provider.secretKey, "storage-secret-key", "", "Cloud storage secret key")
	flag.StringVar(&provider.region, "storage-region", "", "Cloud storage region")
	flag.StringVar(&provider.bucketName, "storage-bucket", "", "Cloud storage bucket name")
	flag.StringVar(&provider.endPoint, "storage-end-point", "", "Cloud storage end point")
	flag.StringVar(&provider.domain, "storage-domain", "", "Cloud storage domain")
}

func (provider *s3Provider) Configure() error {
	provider.logger = logger.GetCurrent().GetLogger(provider.Name())
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(provider.region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(provider.accessKey, provider.secretKey, "")),
	)

	if err != nil {
		provider.logger.Fatalln("Cannot setup storage provider : ", err)
	}

	provider.client = s3.NewFromConfig(cfg)

	provider.logger.Info("Setup storage provider : ", provider.Name())
	return nil
}

func (provider *s3Provider) Run() error {
	return provider.Configure()
}

func (provider *s3Provider) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()

	return c
}

func (provider *s3Provider) SaveUploadedFile(ctx context.Context, data []byte, dst string, contentType string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	params := &s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		Body:        fileBytes,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACL(types.BucketCannedACLPrivate),
	}

	_, err := provider.client.PutObject(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}

func (provider *s3Provider) RemoveUploadedFile(ctx context.Context, dst string) error {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(provider.bucketName),
		Key:    aws.String(dst),
	}
	_, err := provider.client.DeleteObject(context.TODO(), params)
	if err != nil {
		return err
	}

	return nil
}
