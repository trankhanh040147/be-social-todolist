package biz

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	store     ListItemStorage
	requester common.Requester
}

func NewListItemBiz(store ListItemStorage, requester common.Requester) *listItemBiz {
	return &listItemBiz{store: store, requester: requester}
}

func (biz *listItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {
	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester)

	// data, err := biz.store.ListItem(ctx, filter, paging)
	data, err := biz.store.ListItem(ctxStore, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
