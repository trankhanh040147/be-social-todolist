package storage

import (
	"context"
	"go-200lab-g09/module/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
