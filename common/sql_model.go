package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeID    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

// fea_FakeID

func (sql *SQLModel) Mask(dbType DbType) {
	uid := NewUID(uint32(sql.Id), int(dbType), 1)
	sql.FakeID = &uid
}
