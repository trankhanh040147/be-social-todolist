package storage

import (
	"context"
	"go-200lab-g09/module/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, updateData *model.TodoItemUpdate) error {

	if err := s.db.Where(cond).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}
