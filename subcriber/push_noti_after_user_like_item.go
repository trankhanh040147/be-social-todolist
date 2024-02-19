package subscriber

import (
	"context"
	"log"
	"social-todo-list/pubsub"

	goservice "github.com/200Lab-Education/go-sdk"
)

type HasUserId interface {
	GetUserId() int
}

func PushNotiAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Name: "Push notification after user likes item",
		Hdl: func(ctx context.Context, msg *pubsub.Message) error {

			data := msg.Data().(HasUserId)
			log.Println("Push notification to user: ", data.GetUserId())

			return nil
		},
	}
}
