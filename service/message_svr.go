package service

import (
	"fmt"
	"git.ont.io/ontid/otf/message"
)

// MsgService is basic message service implementation
type MsgService struct {
	msgQueue chan message.Message
	quitC    chan struct{}
}

func NewMessageService() *MsgService {
	return &MsgService{
		msgQueue: make(chan message.Message, 64),
		quitC:    make(chan struct{}),
	}
}

func (m *MsgService) HandleOutBound(msg message.Message) error {
	go m.pushMessage(msg)
	return nil
}

func (m *MsgService) pushMessage(msg message.Message) {
	m.msgQueue <- msg
}

func (m *MsgService) popMessage() {
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
