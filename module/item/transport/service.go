package transport

import (
	"context"
	"go-200lab-g09/module/item/model"
)

type ItemUsecase interface {
	CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error
	GetItemById(ctx context.Context, id int) (*model.TodoItem, error)
	UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error
}

type itemService struct {
	useCase ItemUsecase
}
