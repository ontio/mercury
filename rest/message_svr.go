package rest

import (
	"fmt"
	"git.ont.io/ontid/otf/message"
)

// msgService is basic message service implementation
type msgService struct {
	msgQueue chan message.Message
	quitC    chan struct{}
}

func NewMessageService() *msgService {
	return &msgService{
		msgQueue: make(chan message.Message, 64),
		quitC:    make(chan struct{}),
	}
}

func (m *msgService) HandleOutBound(msg message.Message) error {
	go m.pushMessage(msg)
	return nil
}

func (m *msgService) pushMessage(msg message.Message) {
	m.msgQueue <- msg
}

func (m *msgService) popMessage() {
	for {
		select {
		case msg := <-m.msgQueue:
			fmt.Println(":", msg)
		//todo send msg to other agent
		case <-m.quitC:
			return
		}
	}
}
