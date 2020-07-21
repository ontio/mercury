package common

import (
	"encoding/json"
	"net/http"

	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/common/packager"
	"git.ont.io/ontid/otf/common/packager/ecdsa"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	sdk "github.com/ontio/ontology-go-sdk"
)

// MsgService is basic message service implementation
type MsgService struct {
	msgQueue      chan OutboundMsg
	client        *http.Client
	quitC         chan struct{}
	v             vdri.VDRI
	packager      *ecdsa.Packager
	enableEnvelop bool
}

type OutboundMsg struct {
	Msg  Message
	Conn message.Connection
}

func NewMessageService(v vdri.VDRI, ontSdk *sdk.OntologySdk, acct *sdk.Account, enableEnvelop bool) *MsgService {
	ms := &MsgService{
		msgQueue:      make(chan OutboundMsg, 64),
		client:        utils.NewClient(),
		quitC:         make(chan struct{}),
		v:             v,
		packager:      ecdsa.New(ontSdk, acct),
		enableEnvelop: enableEnvelop,
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
		log.Errorf("error on sendmsg:%s\n", err.Error())
	}
	var data []byte
	data, err = json.Marshal(msg.Msg.Content)
	if err != nil {
		log.Errorf("err while sendmsg:%s\n", err)
		return
	}
	if m.enableEnvelop {
		var routerDid string
		if msg.Conn.TheirRouter == nil || len(msg.Conn.TheirRouter) == 0 {
			routerDid = msg.Conn.TheirDid
		} else {
			routerDid = utils.CutDId(msg.Conn.TheirRouter[0])
		}
		msg := &packager.Envelope{
			Message: &packager.MessageData{
				Data:    data,
				MsgType: int(msg.Msg.MessageType),
			},
			FromDID: msg.Conn.MyDid,
			ToDID:   routerDid,
		}
		data, err = m.packager.PackMessage(msg)
		if err != nil {
			log.Errorf("err while sendmsg:%s\n", err)
			return
		}
	}
	log.Infof("url:%s,data:%s\n", url, data)
	_, err = utils.HttpPostData(m.client, url, string(data))
	if err != nil {
		log.Errorf("SendMsg msg url:%s,type:%d,err:%s", url, msg.Msg.MessageType, err)
	}
}

func (m *MsgService) GetServiceURL(msg OutboundMsg) (string, error) {
	var routerDid string
	if msg.Conn.TheirRouter == nil || len(msg.Conn.TheirRouter) == 0 {
		routerDid = msg.Conn.TheirDid
	} else {
		routerDid = msg.Conn.TheirRouter[0]
	}
	doc, err := m.v.GetDIDDoc(utils.CutDId(routerDid))
	if err != nil {
		return "", err
	}
	endpoint, err := doc.GetServicePoint(routerDid)
	if err != nil {
		return "", err
	}
	return endpoint + GetApiName(msg.Msg.MessageType), nil
}
