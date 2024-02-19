package subscriber

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/storage"
	"social-todo-list/pubsub"

	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"
)

type HasItemId interface {
	GetItemId() int
}

func IncreaseLikedCountAfterUserLikeItem(serviceCtx goservice.ServiceContext) subJob {
	return subJob{
		Name: "Increase liked count after user likes item",
		Hdl: func(ctx context.Context, msg *pubsub.Message) error {
			db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
			data := msg.Data().(HasItemId)

			return storage.NewSQLStore(db).IncreaseLikedCount(ctx, data.GetItemId())
		},
	}
}
