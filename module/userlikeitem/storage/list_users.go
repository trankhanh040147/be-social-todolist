package storage

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/model"
)

func (store *sqlStore) ListUsers(ctx context.Context, itemId int, paging *common.Paging) ([]common.SimpleUser, error) {
	var result []model.Like
	db := store.db.Table(model.Like{}.TableName()).Where("item_id = ?", itemId)

	if err := db.Select("user_id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Preload("User").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))
	for index := range users {
		if result[index].User != nil {
			users[index] = *result[index].User
			users[index].UpdatedAt = nil
			users[index].CreatedAt = result[index].CreatedAt
		}
	}

	return users, nil
}
