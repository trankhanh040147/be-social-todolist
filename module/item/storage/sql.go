package storage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func newSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
