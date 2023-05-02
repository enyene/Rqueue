package queue

import (
	"github.com/enyene/Rqueue/messages"
)

type Queue struct {
	Name          string
	InputChannel  chan *messages.Message
	OutputChannel chan *messages.Message
	Head          *queueMessage
	Tail          *queueMessage
	Length        int
}

type queueMessage struct {
	next *queueMessage
	data *messages.Message
}

func NewQueueMesaage(message *messages.Message) *queueMessage {
	return &queueMessage{
		data: message,
	}
}

func NewQueue(name string) *Queue {
	q := Queue{}
	q.InputChannel = make(chan *messages.Message)
	q.OutputChannel = make(chan *messages.Message)

	go q.run()

	return &q
}

func (q *Queue) run() {
run:
	for {
		if q.Head == nil {
			msg, ok := <-q.InputChannel
			msgqueue := NewQueueMesaage(msg)
			if !ok {
				break run

			}
			q.Head = NewQueueMesaage()

		}
	}
}
