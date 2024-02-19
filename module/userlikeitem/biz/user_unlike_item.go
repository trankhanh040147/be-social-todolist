package biz

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
	"social-todo-list/pubsub"
)

type UserUnlikeItemStorage interface {
	Find(ctx context.Context, userId, itemId int) (*model.Like, error)
	Delete(ctx context.Context, userId, itemId int) error
}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStorage
	ps    pubsub.PubSub
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStorage, ps pubsub.PubSub) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store, ps: ps}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)
	if err == common.RecordNotFound {
		return model.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return model.ErrCannotUnlikeItem(err)
	}

	msg := pubsub.NewMessage(&model.Like{
		UserId: userId,
		ItemId: itemId,
	})

	if err := biz.ps.Publish(ctx, common.TopicUserUnlikedItem, msg); err != nil {
		log.Println(err)
	}

	return nil
}
