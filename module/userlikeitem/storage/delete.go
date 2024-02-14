package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
)

func (store *sqlStore) Delete(ctx context.Context, userId, itemId int) error {
	if err := store.db.Table(model.Like{}.TableName()).
		Where("item_id = ? and user_id = ?", itemId, userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
