package messages

import "time"

type Message struct {
	Body *[]byte
	Head *map[string]string
	Time time.Time
}

func NewMessage(head *map[string]string, body *[]byte) *Message {
	return &Message{
		Head: head,
		Body: body,
		Time: time.Now(),
	}
}

func NoHeaderMessage(body *[]byte) *Message {
	head := make(map[string]string)
	return &Message{
		Head: &head,
		Body: body,
		Time: time.Now(),
	}
}
