package storage

import (
	"context"
	"social-todo-list/common"
)

func (store *sqlStore) ListImages(
	ctx context.Context,
	ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	var result []common.Image

	if err := store.db.
		Table(common.Image{}.TableName()).
		Where("id in ?", ids).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
