package storage

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/model"
)

func (store *sqlStore) Create(ctx context.Context, data *model.Like) error {
	if err := store.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
