package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
)

func (store *sqlStore) Create(ctx context.Context, data *model.Like) error {
	if err := store.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
