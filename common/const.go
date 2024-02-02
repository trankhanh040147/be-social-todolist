package common

import "fmt"

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
