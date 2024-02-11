package biz

import (
	"context"
	"errors"
	"go-200lab-g09/common"
	"go-200lab-g09/module/userlikeitem/model"
)

type UserUnlikeItemStorage interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStorage
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStorage) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)

	//if err == common.RecordNotFound {
	if errors.Is(err, common.RecordNotFound) {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	return nil
}
