package storage

import (
	"context"
	"go.opencensus.io/trace"
	"social-todo-list/common"
	"social-todo-list/module/item/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*model.TodoItem, error) {
	_, span := trace.StartSpan(ctx, "item.storage.get")
	defer span.End()

	var data *model.TodoItem

	for _, value := range moreKeys {
		s.db = s.db.Preload(value)
	}

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return data, nil
}
