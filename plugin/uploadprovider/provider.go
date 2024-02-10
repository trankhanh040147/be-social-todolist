package uploadprovider

import (
	"context"

	"go-200lab-g09/common"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, data []byte, dst string, contentType string) (*common.Image, error)
	RemoveUploadedFile(ctx context.Context, dst string) error
}
