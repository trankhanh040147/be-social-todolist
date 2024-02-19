package pubsub

import (
	"context"
	"log"
	"social-todo-list/common"
	"sync"
)

// In-memory
// Buffer channel as queue
// Transmission of messages with specific topic to all subscribers within a group
type localPubSub struct {
	name         string
	messageQueue chan *Message
	mapTopic     map[Topic][]chan *Message
	locker       *sync.RWMutex
}

func NewPubSub(name string) *localPubSub {
	return &localPubSub{
		name:         name,
		messageQueue: make(chan *Message, 10000),
		mapTopic:     make(map[Topic][]chan *Message),
		locker:       new(sync.RWMutex),
	}
}

func (ps *localPubSub) Publish(ctx context.Context, topic Topic, data *Message) error {
	data.SetTopic(topic)

	go func() {
		defer common.Recovery()

		ps.messageQueue <- data
		log.Println("New message published :", data.String())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic Topic) (ch <-chan *Message, unsubscribe func()) {
	c := make(chan *Message)

	ps.locker.Lock()

	if val, ok := ps.mapTopic[topic]; ok {
		val = append(ps.mapTopic[topic], c)
		ps.mapTopic[topic] = val
	} else {
		ps.mapTopic[topic] = []chan *Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapTopic[topic]; ok {
			for index := range chans {
				if chans[index] == c {
					chans = append(chans[:index], chans[index+1:]...)

					ps.locker.Lock()
					ps.mapTopic[topic] = chans
					ps.locker.Unlock()
				}
			}
		}
	}
}

// Send message from message queue to subscribed topic channels
func (ps *localPubSub) run() error {
	go func() {
		defer common.Recovery()

		for {
			msg := <-ps.messageQueue
			log.Println("Message dequeue :", msg.String())

			ps.locker.Lock()

			if subs, ok := ps.mapTopic[msg.Topic()]; ok {
				for index := range subs {
					go func(c chan *Message) {
						defer common.Recovery()
						c <- msg
					}(subs[index])
				}
			}

			ps.locker.Unlock()
		}
	}()

	return nil
}

func (ps *localPubSub) GetPrefix() string {
	return ps.name
}

func (ps *localPubSub) Get() interface{} {
	return ps
}

func (ps *localPubSub) Name() string {
	return ps.name
}

func (*localPubSub) InitFlags() {
}

func (*localPubSub) Configure() error {
	return nil
}

func (ps *localPubSub) Run() error {
	return ps.run()
}

func (*localPubSub) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()

	return c
}
