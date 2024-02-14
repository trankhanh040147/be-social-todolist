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
)

type r2Provider struct {
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

func NewR2Provider(name string) *r2Provider {
	return &r2Provider{
		name: name,
	}
}

func (provider *r2Provider) GetPrefix() string {
	return provider.Name()
}

func (provider *r2Provider) Get() interface{} {
	return provider
}

func (provider *r2Provider) Name() string {
	return provider.name
}

func (provider *r2Provider) InitFlags() {
	flag.StringVar(&provider.accessKey, "storage-access-key", "", "Cloud storage access key")
	flag.StringVar(&provider.secretKey, "storage-secret-key", "", "Cloud storage secret key")
	flag.StringVar(&provider.region, "storage-region", "", "Cloud storage region")
	flag.StringVar(&provider.bucketName, "storage-bucket", "", "Cloud storage bucket name")
	flag.StringVar(&provider.endPoint, "storage-end-point", "", "Cloud storage end point")
	flag.StringVar(&provider.domain, "storage-domain", "", "Cloud storage domain")
}

func (provider *r2Provider) Configure() error {
	provider.logger = logger.GetCurrent().GetLogger(provider.Name())
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: provider.endPoint,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(provider.accessKey, provider.secretKey, "")),
	)

	if err != nil {
		provider.logger.Fatalln("Cannot setup storage provider : ", err)
	}

	provider.client = s3.NewFromConfig(cfg)

	provider.logger.Info("Setup storage provider : ", provider.Name())
	return nil
}

func (provider *r2Provider) Run() error {
	return provider.Configure()
}

func (provider *r2Provider) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()

	return c
}

func (provider *r2Provider) SaveUploadedFile(ctx context.Context, data []byte, dst string, contentType string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	params := &s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		Body:        fileBytes,
		ContentType: aws.String(contentType),
	}

	_, err := provider.client.PutObject(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "r2",
	}

	return img, nil
}

func (provider *r2Provider) RemoveUploadedFile(ctx context.Context, dst string) error {
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
