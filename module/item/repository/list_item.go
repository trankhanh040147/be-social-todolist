package repository

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type ItemLikeStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listItemRepo struct {
	store     ListItemStorage
	likeStore ItemLikeStorage
	requester common.Requester
}

func NewListItemRepo(store ListItemStorage, likeStore ItemLikeStorage, requester common.Requester) *listItemRepo {
	return &listItemRepo{store: store, likeStore: likeStore, requester: requester}
}

func (repo *listItemRepo) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	newCtx := context.WithValue(ctx, common.CurrentUser, repo.requester)

	data, err := repo.store.ListItem(newCtx, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	if len(data) == 0 {
		return data, nil
	}

	ids := make([]int, len(data))
	for index := range ids {
		ids[index] = data[index].Id
	}

	itemLikesMap, err := repo.likeStore.GetItemLikes(newCtx, ids)
	if err != nil {
		return data, nil
	}

	for index := range data {
		data[index].LikedCount = itemLikesMap[data[index].Id]
	}

	return data, nil
}
