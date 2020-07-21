package common

import (
	"encoding/json"
	"git.ont.io/ontid/otf/common/config"
	"net/http"
	"strings"

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
	Cfg           *config.Cfg

}

type OutboundMsg struct {
	Msg  Message
	Conn message.Connection
}

func NewMessageService(v vdri.VDRI, ontSdk *sdk.OntologySdk, acct *sdk.Account, enableEnvelop bool,conf *config.Cfg) *MsgService {
	ms := &MsgService{
		msgQueue:      make(chan OutboundMsg, 64),
		client:        utils.NewClient(),
		quitC:         make(chan struct{}),
		v:             v,
		packager:      ecdsa.New(ontSdk, acct),
		enableEnvelop: enableEnvelop,
		Cfg :conf,
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

	//1. resolve the router
	conn := msg.Conn
	routerlist := MergeRouter(conn.MyRouter, conn.TheirRouter)
	nextrouter, err := m.GetNextRouter(routerlist)
	if err != nil {
		log.Errorf("error on sendmsg:%s\n", err.Error())
		return
	}

	//2. check need forward message
	//f := m.NeedForwardMsg(nextrouter, routerlist)
	var mdata []byte
	var url string
	//if f {
	url, err = m.GetServiceURLByRouter(nextrouter, msg.Msg.MessageType)
	if err != nil {
		log.Errorf("error on sendmsg:%s\n", err.Error())
		return
	}

	mdata, err = json.Marshal(msg.Msg.Content)
	if err != nil {
		log.Errorf("err while sendmsg:%s\n", err)
		return
	}
	var reqbody []byte
	if m.enableEnvelop {
		msg := &packager.Envelope{
			Message: &packager.MessageData{
				Data:    mdata,
				MsgType: int(msg.Msg.MessageType),
			},
			FromDID: m.Cfg.SelfDID,
			ToDID:   utils.CutDId(nextrouter),
		}
		reqbody, err = m.packager.PackMessage(msg)
		if err != nil {
			log.Errorf("err while sendmsg:%s\n", err)
			return
		}
	} else {
		reqbody = mdata
	}

	log.Infof("url:%s,data:%s\n", url, reqbody)
	_, err = utils.HttpPostData(m.client, url, string(reqbody))
	if err != nil {
		log.Errorf("SendMsg msg url:%s,type:%d,err:%s", url, msg.Msg.MessageType, err)
	}

	/*url, err := m.GetServiceURL(msg)
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
	}*/
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
func (m *MsgService) GetServiceURLByRouter(router string, msgType MessageType) (string, error) {
	doc, err := m.v.GetDIDDoc(utils.CutDId(router))
	if err != nil {
		return "", err
	}
	endpoint, err := doc.GetServicePoint(router)
	if err != nil {
		return "", err
	}
	return endpoint + GetApiName(msgType), nil
}

func (m *MsgService) GetNextRouter(routers []string) (string, error) {
	mydid := m.Cfg.SelfDID
	//if the last one is myself
	if strings.EqualFold(mydid, utils.CutDId(routers[len(routers)-1])) {
		return routers[len(routers)-1], nil
	}

	idx, err := RouterLastIndexOf(mydid, routers)
	if err != nil {
		log.Errorf("error on sendmsg:%s\n", err.Error())
		return "", err
	}
	if strings.EqualFold(mydid, utils.CutDId(routers[idx+1])) {
		return routers[idx+2], nil
	}
	return routers[idx+1], nil
}

func (m *MsgService) NeedForwardMsg(router string, routers []string) bool {
	return strings.EqualFold(router, routers[len(routers)-1])
}