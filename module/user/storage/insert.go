package storage

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/module/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.UserCreate) error {
	db := s.db.Begin()
	// data.PrepareForInsert()

	// if has db.Begin() then must has db.Rollback() or db.Commit() to end transaction, otherwise it will be locked

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
