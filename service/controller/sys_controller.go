package controller

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/vdri"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	InvitationKey    = "Invitation"
	ConnectionReqKey = "ConnectionReq"
	ConnectionKey    = "Connection"

	ACK_SUCCEED = "succeed"
	ACK_FAILED  = "failed"
)

type Syscontroller struct {
	account *sdk.Account
	did     vdri.Did
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
}

func NewSyscontroller(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService, did vdri.Did) Syscontroller {
	s := Syscontroller{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
	}
	s.Initiate(nil)
	return s
}

func (s Syscontroller) Name() string {
	return "syscontroller"
}

func (s Syscontroller) Initiate(param service.ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s Syscontroller) Process(msg message.Message) (service.ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	//for system
	case message.InvitationType:
		fmt.Println("resolve invitation")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		//todo verify request
		invitation, ok := msg.Content.(*message.Invitation)
		if !ok {
			return nil, fmt.Errorf("message format is not correct")
		}
		//set uuid to invitation id
		invitation.Id = uuid.New().String()

		//store the invitation
		err := s.SaveInvitation(*invitation)
		if err != nil {
			return nil, err
		}

		return service.ServiceResp{
			Message: invitation,
		}, nil

	//not in use anymore
	//case message.SendConnectionRequestType:
	//	//send connection req for agent
	//	fmt.Println("resolve send connection request")
	//	//todo verify request
	//	if msg.Content == nil {
	//		return nil, fmt.Errorf("message content is nil")
	//	}
	//	cr := msg.Content.(*message.ConnectionRequest)
	//	cr.Id = uuid.New().String()
	//	err := s.SaveConnectionRequest(*cr, ConnectionRequestSent)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	//send the connection req to target service endpoint
	//	msg.Content = cr
	//	msg.MessageType = message.ConnectionRequestType
	//	jsonbytes,err := json.Marshal(msg)
	//	if err != nil{
	//		return nil,err
	//	}
	//	msg.JsonBytes = jsonbytes
	//
	//	err = s.msgsvr.HandleOutBound(msg)
	//	if err != nil {
	//		return nil, err
	//	}
	//	middleware.Log.Infof("SendConnectionReq:%v", cr)
	//	//no need to pass incoming param
	//	return nil, nil

	case message.ConnectionRequestType:
		fmt.Println("resolve connection request")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionRequest)
		//ivid := req.Thread.ID
		ivrc, err := s.GetInvitation(req.InvitationId)
		if err != nil {
			fmt.Printf("err on GetInvitation:%s\n", err.Error())
			return nil, err
		}

		//update connection to request received state
		err = s.SaveConnectionRequest(*req, service.ConnectionRequestReceived)
		if err != nil {
			fmt.Printf("err on SaveConnectionRequest:%s\n", err.Error())
			return nil, err
		}

		//update invitation to used state
		err = s.UpdateInvitation(ivrc.Invitation.Id, service.InvitationUsed)
		if err != nil {
			fmt.Printf("err on UpdateInvitation:%s\n", err.Error())
			return nil, err
		}

		//send response outbound
		res := new(message.ConnectionResponse)
		res.Id = uuid.New().String()
		res.Thread = message.Thread{
			ID: req.Id,
		}
		//todo define the response type
		res.Type = vdri.ConnectionResponseSpec
		//self conn
		res.Connection = message.Connection{
			MyDid:          ivrc.Invitation.Did,
			MyServiceId:    ivrc.Invitation.ServiceId,
			TheirDid:       req.Connection.MyDid,
			TheirServiceId: req.Connection.MyServiceId,
		}

		outmsg := message.Message{
			MessageType: message.ConnectionResponseType,
			Content:     res,
		}
		err = s.msgsvr.HandleOutBound(service.OutboundMsg{
			Msg:  outmsg,
			Conn: res.Connection,
		})
		if err != nil {
			if err != nil {
				fmt.Printf("err on HandleOutBound:%s\n", err.Error())
				return nil, err
			}
			return nil, err
		}
		return nil, nil

	case message.ConnectionResponseType:
		fmt.Println("resolve connection response")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionResponse)
		connid := req.Thread.ID

		//2. create and save a connection object
		err := s.SaveConnection(req.Connection.TheirDid,
			req.Connection.TheirServiceId,
			req.Connection.MyDid,
			req.Connection.MyServiceId)
		if err != nil {
			fmt.Printf("err on SaveConnection:%s\n", err.Error())
			return nil, err
		}

		//3. send ACK back
		ack := message.ConnectionACK{
			Type:       vdri.ConnectionACKSpec,
			Id:         uuid.New().String(),
			Thread:     message.Thread{ID: connid},
			Status:     ACK_SUCCEED,
			Connection: service.ReverseConnection(req.Connection),
		}

		outmsg := message.Message{
			MessageType: message.ConnectionACKType,
			Content:     ack,
		}
		err = s.msgsvr.HandleOutBound(service.OutboundMsg{
			Msg:  outmsg,
			Conn: ack.Connection,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	case message.ConnectionACKType:
		fmt.Println("resolve ConnectionACK")
		req := msg.Content.(*message.ConnectionACK)
		//1. update connection request to receive ack state
		if req.Status != ACK_SUCCEED {
			//todo remove connectionreq when failed?
			return nil, fmt.Errorf("got failed ACK ")
		}
		connid := req.Thread.ID
		err := s.UpdateConnectionRequest(connid, service.ConnectionACKReceived)
		if err != nil {
			fmt.Printf("err on UpdateConnectionRequest:%s\n", err.Error())
			return nil, err
		}
		//2. create and save a connection object
		cr, err := s.GetConnectionRequest(connid)
		if err != nil {
			fmt.Printf("err on GetConnectionRequest:%s\n", err.Error())
			return nil, err
		}

		err = s.SaveConnection(cr.ConnReq.Connection.TheirDid,
			cr.ConnReq.Connection.TheirServiceId,
			cr.ConnReq.Connection.MyDid,
			cr.ConnReq.Connection.MyServiceId)
		if err != nil {
			fmt.Printf("err on SaveConnection:%s\n", err.Error())
			return nil, err
		}
		return nil, nil

	case message.SendGeneralMsgType:
		fmt.Println("resolve SendGeneralMsgType")
		req := msg.Content.(*message.BasicMessage)
		data, err := json.Marshal(req)
		if err != nil {
			fmt.Printf("err on Marshal:%s\n", err.Error())
			return nil, err
		}
		fmt.Println("we got a message: %s", data)

		conn, err := s.GetConnection(req.Connection.MyDid, req.Connection.TheirDid, req.Connection.TheirServiceId)
		if err != nil {
			fmt.Printf("err on GetConnection:%s\n", err.Error())
			return nil, err
		}
		req.Type = vdri.BasicMsgSpec
		req.Id = uuid.New().String()

		om := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.SendGeneralMsgType,
				Content:     req,
			},
			Conn: conn,
		}
		err = s.msgsvr.HandleOutBound(om)
		if err != nil {
			fmt.Printf("err on HandleOutBound:%s\n", err.Error())
			return nil, err
		}
		return nil, nil

	case message.ReceiveGeneralMsgType:
		fmt.Println("resolve ReceiveGeneralMsgType")
		req := msg.Content.(*message.BasicMessage)
		data, err := json.Marshal(req)
		if err != nil {
			fmt.Printf("err on Marshal:%s\n", err.Error())
			return nil, err
		}
		fmt.Printf("we got a message: %s\n", data)
		return nil, nil

	default:
		return service.Skipmessage(msg)
	}

	resp := service.ServiceResp{}
	return resp, nil
}
func (s Syscontroller) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}

func (s Syscontroller) sign(data []byte) ([]byte, error) {
	sig, err := signature.Sign(signature.SHA256withECDSA, s.account.PrivateKey, data, nil)
	if err != nil {
		return nil, err
	}
	return signature.Serialize(sig)
}

func (s Syscontroller) toMap(v interface{}) (map[string]interface{}, error) {
	return structs.Map(v), nil
}

//

func (s Syscontroller) SaveInvitation(iv message.Invitation) error {

	key := fmt.Sprintf("%s_%s", InvitationKey, iv.Id)
	b, err := s.store.Has([]byte(key))
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("invitation with id:%s existed", iv.Id)
	}

	rec := service.InvitationRec{
		Invitation: iv,
		State:      service.InvitationInit,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put([]byte(key), bs)
}

func (s Syscontroller) GetInvitation(id string) (*service.InvitationRec, error) {
	key := []byte(fmt.Sprintf("%s_%s", InvitationKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}

	rec := new(service.InvitationRec)

	err = json.Unmarshal(data, rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (s Syscontroller) UpdateInvitation(id string, state service.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", InvitationKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(service.InvitationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	//fixme introduce some FSM
	if rec.State >= state {
		return fmt.Errorf("error state with id:%s", id)
	}
	rec.State = state
	bts, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}

func (s Syscontroller) SaveConnectionRequest(cr message.ConnectionRequest, state service.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, cr.Id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("connection request with id:%s existed", cr.Id)
	}
	rec := service.ConnectionRequestRec{
		ConnReq: cr,
		State:   state,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put(key, bs)
}

func (s Syscontroller) GetConnectionRequest(id string) (*service.ConnectionRequestRec, error) {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	cr := new(service.ConnectionRequestRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (s Syscontroller) UpdateConnectionRequest(id string, state service.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(service.ConnectionRequestRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}

	if rec.State >= state {
		return fmt.Errorf("error state with id:%s", id)
	}

	rec.State = state
	bts, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}

func (s Syscontroller) SaveConnection(myDID, myServiceId, theirDID, theirServiceID string) error {

	cr := new(service.ConnectionRec)

	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, myDID))
	exist, err := s.store.Has(key)
	if err != nil {
		return err
	}

	if exist {
		data, err := s.store.Get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, cr)
		if err != nil {
			return err
		}
		cr.Connections[fmt.Sprintf("%s_%s", theirDID, theirServiceID)] = message.Connection{
			MyDid:          myDID,
			MyServiceId:    myServiceId,
			TheirDid:       theirDID,
			TheirServiceId: theirServiceID,
		}
	} else {
		cr.OwnerDID = myDID
		m := make(map[string]message.Connection)
		m[fmt.Sprintf("%s_%s", theirDID, theirServiceID)] = message.Connection{
			TheirDid:       theirDID,
			TheirServiceId: theirServiceID,
		}
		cr.Connections = m
	}
	bts, err := json.Marshal(cr)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}

func (s Syscontroller) GetConnection(myDID, theirDID, theirServiceID string) (message.Connection, error) {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, myDID))
	data, err := s.store.Get(key)
	if err != nil {
		return message.Connection{}, err
	}
	cr := new(service.ConnectionRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return message.Connection{}, err
	}
	c, ok := cr.Connections[fmt.Sprintf("%s_%s", theirDID, theirServiceID)]
	if !ok {
		return message.Connection{}, fmt.Errorf("connection not found!")
	}

	return c, nil
}
