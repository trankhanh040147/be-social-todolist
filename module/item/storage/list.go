package storage

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {

	var result []model.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	// Get items of requester only
	requester := ctx.Value(common.CurrentUser).(common.Requester)
	db.Where("user_id = ?", requester.GetUserId())

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Select("id").Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Select("*").Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {

		return nil, err
	}

	return result, nil
}
