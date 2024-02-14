package biz

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
)

type ListItemRepo interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	repo      ListItemRepo
	requester common.Requester
}

func NewListItemBiz(repo ListItemRepo, requester common.Requester) *listItemBiz {
	return &listItemBiz{repo: repo, requester: requester}
}

func (biz *listItemBiz) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, biz.requester)

	data, err := biz.repo.ListItem(newCtx, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
