package biz

import (
	"context"
	"go-200lab-g09/module/userlikeitem/model"
)

type UserLikeItemStorage interface {
	Create(ctx context.Context, data *model.Like) error
}

type userLikeItemBiz struct {
	store UserLikeItemStorage
}

func NewUserLikeItemBiz(store UserLikeItemStorage) *userLikeItemBiz {
	return &userLikeItemBiz{store: store}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	return nil
}
