package pubsub

import (
	"fmt"
	"time"
)

type Topic string

type Message struct {
	id        string
	topic     Topic // ignorable
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message %s value %v", m.topic, m.data)
}

func (m *Message) Topic() Topic {
	return m.topic
}

func (m *Message) SetTopic(topic Topic) {
	m.topic = topic
}

func (m *Message) Data() interface{} {
	return m.data
}
