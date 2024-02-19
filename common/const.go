package common

import (
	"fmt"
)

type DbType int

const (
	DbTypeItem DbType = 1
	DbTypeUser DbType = 2
)

const (
	PluginDBMain  = "mysql"
	PluginJWT     = "jwt"
	PluginR2      = "r2"
	PluginS3      = "s3"
	CurrentUser   = "current_user"
	PluginPubSub  = "pubsub"
	PluginItemAPI = "item-api"
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
