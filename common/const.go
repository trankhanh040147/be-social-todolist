package common

import (
	"fmt"
)

const CurrentUser = "current_user"

type DbType int

const (
	DbTypeItem DbType = 1
	DbTypeUser DbType = 2
)

const (
	PluginDBMain = "mysql"
	PluginJWT    = "jwt"
	PluginR2     = "r2"
	PluginS3     = "s3"
	PluginPubSub = "pubsub"

	TopicUserLikedItem   = "TopicUserLikedItem"
	TopicUserUnlikedItem = "TopicUserUnlikedItem"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

type TokenPayLoad struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayLoad) UserId() int {
	return p.UId
}

func (p TokenPayLoad) Role() string {
	return p.URole
}

type Requester interface {
	GetUserId() int
	GetRole() string
	GetEmail() string
}

func IsAdminOrMod(requester Requester) bool {
	return requester.GetRole() == "admin" || requester.GetRole() == "mod"
}
