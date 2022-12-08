package message

import "time"

type Message struct {
	Text string
	Time time.Time
}

func New(text string) *Message {
	return &Message{
		Text: text,
		Time: time.Now(),
	}
}
