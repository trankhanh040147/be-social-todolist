package pubsub

import (
	"context"
)

type PubSub interface {
	Publish(ctx context.Context, topic Topic, data *Message) error
	Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, unsubscribe func())
}
