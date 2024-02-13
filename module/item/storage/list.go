package storage

import (
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
	"golang.org/x/net/context"
	"strings"
)

func (store *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := store.db.
		Table(model.TodoItem{}.TableName()).
		Where("status <> ?", "Deleted")

	// Get items of requester only
	// requester := ctx.Value(common.CurrentUser).(common.Requester)
	// db = db.Where("user_id = ?", requester.GetUserId())

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, value := range moreKeys {
		db = db.Preload(value)
	}

	if cursor := strings.TrimSpace(paging.FakeCursor); cursor != "" {
		id, err := common.UIDFromBase58(cursor)
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", id.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Select("*").
		Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	size := len(result)
	if size > 0 {
		result[size-1].Mask()
		paging.NextCursor = result[size-1].FakeID.String()
	}

	return result, nil
}
