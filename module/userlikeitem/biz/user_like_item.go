package biz

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/model"
	"log"
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
