package model

import (
	"errors"
	"go-200lab-g09/common"
	"strings"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
)

type TodoItem struct {
	common.SQLModel        // embed struct
	Title           string `json:"title" gorm:"column:title;"`
	Description     string `json:"description" gorm:"column:;description"`
	Status          string `json:"status" gorm:"column:status;"`
}

// >> Why it do not have receiver like (t TodoItem) ? --> it apply for all TodoItem objects
func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:;description"`
}

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	// use pointer(*) to allow update data to "", 0, false... except nil
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:;description"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
