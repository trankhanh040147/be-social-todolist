package storage

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, updateData *model.TodoItemUpdate) error {

	if err := s.db.Where(cond).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}

func (store *sqlStore) IncreaseLikedCount(ctx context.Context, id int) error {
	if err := store.db.Table(model.TodoItem{}.TableName()).
		Where("id = ?", id).
		Update("liked_count ", gorm.Expr("liked_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (store *sqlStore) DecreaseLikedCount(ctx context.Context, id int) error {
	if err := store.db.Table(model.TodoItem{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
