package service

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	"net/http"
	"strings"

	"git.ont.io/ontid/otf/message"
)

// MsgService is basic message service implementation
type MsgService struct {
	msgQueue chan OutboundMsg
	client   *http.Client
	quitC    chan struct{}
	vdri     vdri.VDRI
}

type OutboundMsg struct {
	Msg  message.Message
	Conn message.Connection
}

func NewMessageService(vdri vdri.VDRI) *MsgService {
	ms := &MsgService{
		msgQueue: make(chan OutboundMsg, 64),
		client:   &http.Client{},
		quitC:    make(chan struct{}),
		vdri:     vdri,
	}
	go ms.popMessage()
	return ms
}

func (m *MsgService) HandleOutBound(omsg OutboundMsg) error {
	go m.pushMessage(omsg)
	return nil
}

func (m *MsgService) pushMessage(msg OutboundMsg) {
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

func (m *MsgService) SendMsg(msg OutboundMsg) {
	url, err := m.GetServiceURL(msg)
	if err != nil {
		middleware.Log.Errorf("error on sendmsg:%s\n", err.Error())
	}
	data, err := json.Marshal(msg.Msg.Content)
	if err != nil {
		middleware.Log.Errorf("err while sendmsg:%s\n", err)
		return
	}
	middleware.Log.Infof("url:%s,data:%s\n", url, data)
	err = m.HttpPostData(url, string(data))
	if err != nil {
		middleware.Log.Errorf("SendMsg msg url:%s,type:%d,err:%s", url, msg.Msg.MessageType, err)
	}
}

func (m *MsgService) HttpPostData(url, data string) error {
	_, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post request:%s error:%s", data, err)
	}
	//todo analyze the resp???

	return nil
}

func (m *MsgService) GetServiceURL(msg OutboundMsg) (string, error) {

	doc, err := m.vdri.GetDIDDoc(msg.Conn.TheirDid)
	if err != nil {
		return "", err
	}
	endpoint, err := doc.GetServicePoint(fmt.Sprintf("%s#%s", msg.Conn.TheirDid, msg.Conn.TheirServiceId))
	if err != nil {
		return "", err
	}
	return endpoint + utils.GetApiName(msg.Msg.MessageType), nil

}
