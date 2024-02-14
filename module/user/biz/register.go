package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(storage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: storage, hasher: hasher}
}

func (biz *registerBusiness) Register(ctx context.Context, data *model.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		// status = 0 --> return error user has disabled

		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hardcode

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
