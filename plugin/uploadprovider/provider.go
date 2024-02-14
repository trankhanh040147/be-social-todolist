package uploadprovider

import (
	"context"

	"social-todo-list/common"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, data []byte, dst string, contentType string) (*common.Image, error)
	RemoveUploadedFile(ctx context.Context, dst string) error
}
