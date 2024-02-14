package biz

import (
	"context"
	"errors"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemBiz struct {
	store     UpdateItemStorage
	requester common.Requester
}

func NewUpdateItemBiz(store UpdateItemStorage, requester common.Requester) *updateItemBiz {
	return &updateItemBiz{store: store, requester: requester}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return common.ErrEntityDeleted(model.EntityName)
	}

	isOwner := biz.requester.GetUserId() == data.UserId

	if !isOwner && !common.IsAdminOrMod(biz.requester) {
		return common.ErrNoPermission(errors.New("you don't have permission to update this item"))
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
