package controller

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	"github.com/fatih/structs"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	InvitationKey    = "Invitation"
	ConnectionReqKey = "ConnectionReq"
	ConnectionKey    = "Connection"
	GeneralMsgKey    = "General"

	ACK_SUCCEED = "succeed"
	ACK_FAILED  = "failed"
)

type Syscontroller struct {
	account *sdk.Account
	//did     vdri.Did
	cfg    *config.Cfg
	store  store.Store
	msgsvr *service.MsgService
}

func NewSyscontroller(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService) Syscontroller {
	s := Syscontroller{
		account: acct,
		//did:     did,
		cfg:    cfg,
		store:  db,
		msgsvr: msgsvr,
	}
	err := s.Initiate(nil)
	if err != nil {
		panic(err)
	}
	return s
}

func (s Syscontroller) Name() string {
	return "syscontroller"
}

func (s Syscontroller) Initiate(param service.ParameterInf) error {
	middleware.Log.Infof("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s Syscontroller) Process(msg message.Message) (service.ControllerResp, error) {
	middleware.Log.Infof("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	//for system
	case message.InvitationType:
		middleware.Log.Infof("resolve invitation")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		//todo verify request
		invitation, ok := msg.Content.(*message.Invitation)
		if !ok {
			return nil, fmt.Errorf("message format is not correct")
		}

		//store the invitation
		err := s.SaveInvitation(*invitation)
		if err != nil {
			return nil, err
		}

		return service.ServiceResponse{
			Message: invitation,
		}, nil

	case message.ConnectionRequestType:
		middleware.Log.Infof("resolve connection request")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionRequest)
		//ivid := req.Thread.ID
		ivrc, err := s.GetInvitation(req.Connection.TheirDid, req.InvitationId)
		if err != nil {
			middleware.Log.Infof("err on GetInvitation:%s\n", err.Error())
			return nil, err
		}

		//update connection to request received state
		err = s.SaveConnectionRequest(*req, message.ConnectionRequestReceived)
		if err != nil {
			middleware.Log.Infof("err on SaveConnectionRequest:%s\n", err.Error())
			return nil, err
		}

		//update invitation to used state
		err = s.UpdateInvitation(ivrc.Invitation.Did, ivrc.Invitation.Id, message.InvitationUsed)
		if err != nil {
			middleware.Log.Infof("err on UpdateInvitation:%s\n", err.Error())
			return nil, err
		}

		//send response outbound
		res := new(message.ConnectionResponse)
		res.Id = utils.GenUUID()
		res.Thread = message.Thread{
			ID: req.Id,
		}
		//todo define the response type
		res.Type = vdri.ConnectionResponseSpec
		//self conn
		res.Connection = message.Connection{
			MyDid:          ivrc.Invitation.Did,
			MyServiceId:    ivrc.Invitation.ServiceId,
			MyRouter:       ivrc.Invitation.Router,
			TheirDid:       req.Connection.MyDid,
			TheirServiceId: req.Connection.MyServiceId,
			TheirRouter:    req.Connection.MyRouter,
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
			middleware.Log.Errorf("err on HandleOutBound:%s\n", err.Error())
			return nil, err
		}
		return nil, nil

	case message.ConnectionResponseType:
		middleware.Log.Infof("resolve connection response")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionResponse)
		connid := req.Thread.ID

		//2. create and save a connection object
		err := s.SaveConnection(req.Connection)
		if err != nil {
			middleware.Log.Errorf("err on SaveConnection:%s\n", err.Error())
			return nil, err
		}

		//3. send ACK back
		ack := message.ConnectionACK{
			Type:       vdri.ConnectionACKSpec,
			Id:         utils.GenUUID(),
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
		middleware.Log.Infof("resolve ConnectionACK")
		req := msg.Content.(*message.ConnectionACK)
		//1. update connection request to receive ack state
		if req.Status != ACK_SUCCEED {
			//todo remove connectionreq when failed?
			return nil, fmt.Errorf("got failed ACK ")
		}
		connid := req.Thread.ID
		err := s.UpdateConnectionRequest(req.Connection.TheirDid, connid, message.ConnectionACKReceived)
		if err != nil {
			middleware.Log.Errorf("err on UpdateConnectionRequest:%s\n", err.Error())
			return nil, err
		}
		//2. create and save a connection object
		cr, err := s.GetConnectionRequest(req.Connection.TheirDid, connid)
		if err != nil {
			middleware.Log.Errorf("err on GetConnectionRequest:%s\n", err.Error())
			return nil, err
		}

		err = s.SaveConnection(service.ReverseConnection(cr.ConnReq.Connection))
		if err != nil {
			middleware.Log.Errorf("err on SaveConnection:%s\n", err.Error())
			return nil, err
		}
		return nil, nil

	case message.SendGeneralMsgType:
		middleware.Log.Infof("resolve SendGeneralMsgType")
		req := msg.Content.(*message.BasicMessage)
		data, err := json.Marshal(req)
		if err != nil {
			middleware.Log.Errorf("err on Marshal:%s\n", err.Error())
			return nil, err
		}
		middleware.Log.Infof("we got a message: %s", data)

		conn, err := s.GetConnection(req.Connection.MyDid, req.Connection.TheirDid, req.Connection.TheirServiceId)
		if err != nil {
			middleware.Log.Errorf("err on GetConnection:%s\n", err.Error())
			return nil, err
		}
		req.Type = vdri.BasicMsgSpec
		req.Id = utils.GenUUID()

		om := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.SendGeneralMsgType,
				Content:     req,
			},
			Conn: conn,
		}
		err = s.msgsvr.HandleOutBound(om)
		if err != nil {
			middleware.Log.Errorf("err on HandleOutBound:%s\n", err.Error())
			return nil, err
		}
		return nil, nil

	case message.ReceiveGeneralMsgType:
		middleware.Log.Infof("resolve ReceiveGeneralMsgType")
		req := msg.Content.(*message.BasicMessage)

		//todo remove me
		data, err := json.Marshal(req)
		if err != nil {
			middleware.Log.Errorf("err on Marshal:%s\n", err.Error())
			return nil, err
		}
		middleware.Log.Infof("we got a message: %s\n", data)
		//todo end
		return nil, s.SaveGeneralMsg(req)

	case message.QueryGeneralMessageType:
		middleware.Log.Infof("resolve ReceiveGeneralMsgType")
		req := msg.Content.(*message.QueryGeneralMessageRequest)
		ret, err := s.QueryGeneraMsg(req.DID, req.Latest, req.RemoveAfterRead)
		if err != nil {
			return nil, err
		}
		return service.ServiceResponse{
			Message: ret,
		}, nil

	default:
		return service.Skipmessage(msg)
	}

}
func (s Syscontroller) Shutdown() error {
	middleware.Log.Infof("%s shutdown\n", s.Name())
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

func (s Syscontroller) QueryGeneraMsg(did string, latest bool, removeAfterRead bool) ([]message.BasicMessage, error) {
	key := []byte(fmt.Sprintf("%s_%s", GeneralMsgKey, did))
	b, err := s.store.Has(key)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	}
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	rec := new(message.GeneralMsgRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return nil, err
	}
	if rec.Msglist == nil || len(rec.Msglist) == 0 {
		return nil, nil
	}
	retlist := make([]message.BasicMessage, 0)
	if latest {
		retlist = rec.Msglist[len(rec.Msglist)-1:]
		if removeAfterRead {
			rec.Msglist = rec.Msglist[0 : len(rec.Msglist)-1]
			data, err := json.Marshal(rec)
			if err != nil {
				return nil, err
			}
			err = s.store.Put(key, data)
			if err != nil {
				return nil, err
			}
		}

	} else {
		retlist = rec.Msglist
		if removeAfterRead {
			err = s.store.Delete(key)
			if err != nil {
				return nil, err
			}
		}
	}
	return retlist, nil

}

func (s Syscontroller) SaveGeneralMsg(m *message.BasicMessage) error {
	key := []byte(fmt.Sprintf("%s_%s", GeneralMsgKey, m.Connection.TheirDid))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	rec := new(message.GeneralMsgRec)
	if b {
		data, err := s.store.Get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, rec)
		if err != nil {
			return err
		}
		rec.Msglist = append(rec.Msglist, *m)
	} else {
		rec.Msglist = []message.BasicMessage{*m}
	}

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

//

func (s Syscontroller) SaveInvitation(iv message.Invitation) error {

	key := fmt.Sprintf("%s_%s_%s", InvitationKey, iv.Did, iv.Id)
	b, err := s.store.Has([]byte(key))
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("invitation with id:%s existed", iv.Id)
	}

	rec := message.InvitationRec{
		Invitation: iv,
		State:      message.InvitationInit,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put([]byte(key), bs)
}

func (s Syscontroller) GetInvitation(did, id string) (*message.InvitationRec, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", InvitationKey, did, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}

	rec := new(message.InvitationRec)

	err = json.Unmarshal(data, rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (s Syscontroller) UpdateInvitation(did, id string, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", InvitationKey, did, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(message.InvitationRec)
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

func (s Syscontroller) SaveConnectionRequest(cr message.ConnectionRequest, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", ConnectionReqKey, cr.Connection.TheirDid, cr.Id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("connection request with id:%s existed", cr.Id)
	}
	rec := message.ConnectionRequestRec{
		ConnReq: cr,
		State:   state,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put(key, bs)
}

func (s Syscontroller) GetConnectionRequest(did, id string) (*message.ConnectionRequestRec, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", ConnectionReqKey, did, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	cr := new(message.ConnectionRequestRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (s Syscontroller) UpdateConnectionRequest(did, id string, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", ConnectionReqKey, did, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(message.ConnectionRequestRec)
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

func (s Syscontroller) SaveConnection(con message.Connection) error {

	cr := new(message.ConnectionRec)

	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, con.MyDid))
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
		cr.Connections[fmt.Sprintf("%s_%s", con.TheirDid, con.TheirServiceId)] = con
	} else {
		cr.OwnerDID = con.MyDid
		m := make(map[string]message.Connection)
		m[fmt.Sprintf("%s_%s", con.TheirDid, con.TheirServiceId)] = con
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
	cr := new(message.ConnectionRec)
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
