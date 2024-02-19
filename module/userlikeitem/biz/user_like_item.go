package biz

import (
	"context"
	"log"
	"social-todo-list/common"
	"social-todo-list/module/userlikeitem/model"
	"social-todo-list/pubsub"
)

type UserLikeItemStorage interface {
	Create(ctx context.Context, data *model.Like) error
}

type IncreaseLikedCountStorage interface {
	IncreaseLikedCount(ctx context.Context, id int) error
}

type userLikeItemBiz struct {
	store UserLikeItemStorage
	ps    pubsub.PubSub
}

func NewUserLikeItemBiz(store UserLikeItemStorage, ps pubsub.PubSub) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, ps: ps}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *model.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return model.ErrCannotLikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikedItem, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
