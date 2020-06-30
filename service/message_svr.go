package service

import (
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/utils"
	"net/http"
	"strings"

	"git.ont.io/ontid/otf/message"
)

// MsgService is basic message service implementation
type MsgService struct {
	msgQueue chan message.Message
	client   *http.Client
	quitC    chan struct{}
}

func NewMessageService() *MsgService {
	return &MsgService{
		msgQueue: make(chan message.Message, 64),
		client:   &http.Client{},
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
			m.SendMsg(msg)
		case <-m.quitC:
			return
		}
	}
}

func (m *MsgService) SendMsg(msg message.Message) {
	url := "http://127.0.0.1:8080" + utils.GetApiName(msg.MessageType)
	err := m.HttpPostData(url, string(msg.JsonBytes))
	if err != nil {
		middleware.Log.Errorf("SendMsg msg url:%s,type:%d,err:%s", url, msg.MessageType, err)
	}
}

func (m *MsgService) HttpPostData(url, data string) error {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	if err != nil {
		middleware.Log.Errorf("post data err:%s", err)
		return err
	}
	_, err = m.client.Do(req)
	if err != nil {
		middleware.Log.Errorf("httpPostData do err:%s", err)
		return err
	}
	return nil
}
