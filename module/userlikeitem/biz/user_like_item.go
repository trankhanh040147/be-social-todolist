package biz

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
)

type UserLikeItemStorage interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseLikedCountStorage interface {
	IncreaseLikedCount(ctx context.Context, id int) error
}

type userLikeItemBiz struct {
	store     UserLikeItemStorage
	itemStore IncreaseLikedCountStorage
}

func NewUserLikeItemBiz(store UserLikeItemStorage, itemStore IncreaseLikedCountStorage) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	go func() {
		defer common.Recovery()

		if err := biz.itemStore.IncreaseLikedCount(ctx, data.ItemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
