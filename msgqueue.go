package rqueue

import (
	"github.com/enyene/Rqueue/messages"
	"github.com/enyene/Rqueue/queue"
)

type MessageQueue struct {
	Name                   string
	MessageInput           chan *messages.Message
	MessageOutput          chan *messages.Message
	queue                  *queue.QueueMessage
	subscribers            *string
	running                bool
	messagesSentLastSecond uint64
}
