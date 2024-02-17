package storage

import (
	"context"
	"social-todo-list/common"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	if err := store.db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
