package queue

import (
	"github.com/enyene/Rqueue/messages"
)

type queueMessage struct {
	next *queueMessage
	data *messages.Message
}

func NewQueueMesaage(message *messages.Message) *queueMessage {
	return &queueMessage{
		data: message,
	}
}
