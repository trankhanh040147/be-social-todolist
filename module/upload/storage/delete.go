package storage

import (
	"context"
	"social-todo-list/common"
)

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	if err := store.db.Table(common.Image{}.TableName()).
		Where("id in ?", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
