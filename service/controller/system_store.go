package controller

import (
	"encoding/json"
	"fmt"

	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/utils"
)

func (s *SystemController) SaveInvitation(iv message.Invitation) error {
	key := fmt.Sprintf("%s_%s_%s", utils.InvitationKey, iv.Did, iv.Id)
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

func (s SystemController) GetInvitation(did, id string) (*message.InvitationRec, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", utils.InvitationKey, did, id))
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

func (s SystemController) UpdateInvitation(did, id string, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", utils.InvitationKey, did, id))
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

func (s SystemController) SaveConnectionRequest(cr message.ConnectionRequest, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", utils.ConnectionReqKey, cr.Connection.TheirDid, cr.Id))
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

func (s SystemController) GetConnectionRequest(did, id string) (*message.ConnectionRequestRec, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", utils.ConnectionReqKey, did, id))
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

func (s *SystemController) SaveConnection(con message.Connection) error {
	cr := new(message.ConnectionRec)

	key := []byte(fmt.Sprintf("%s_%s", utils.ConnectionKey, con.MyDid))
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
		cr.Connections[con.TheirDid] = con
	} else {
		cr.OwnerDID = con.MyDid
		m := make(map[string]message.Connection)
		m[con.TheirDid] = con
		cr.Connections = m
	}
	bts, err := json.Marshal(cr)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}

func (s *SystemController) UpdateConnectionRequest(did, id string, state message.ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", utils.ConnectionReqKey, did, id))
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

func (s *SystemController) GetConnection(myDID, theirDID string) (message.Connection, error) {
	key := []byte(fmt.Sprintf("%s_%s", utils.ConnectionKey, myDID))
	data, err := s.store.Get(key)
	if err != nil {
		return message.Connection{}, err
	}
	cr := new(message.ConnectionRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return message.Connection{}, err
	}
	c, ok := cr.Connections[theirDID]
	if !ok {
		return message.Connection{}, fmt.Errorf("connection not found!")
	}

	return c, nil
}

func (s *SystemController) DeleteConnection(myDID, theirDID string) error {
	key := []byte(fmt.Sprintf("%s_%s", utils.ConnectionKey, myDID))

	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	cr := new(message.ConnectionRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return err
	}
	delete(cr.Connections, theirDID)
	data, err = json.Marshal(cr)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

func (s *SystemController) SaveGeneralMsg(m *message.BasicMessage, send bool) error {
	var did string
	if send {
		did = m.Connection.MyDid
	} else {
		did = m.Connection.TheirDid
	}
	key := []byte(fmt.Sprintf("%s_%s", utils.GeneralMsgKey, did))
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

func (s *SystemController) QueryGeneraMsg(did string, latest bool, removeAfterRead bool) ([]message.BasicMessage, error) {
	key := []byte(fmt.Sprintf("%s_%s", utils.GeneralMsgKey, did))
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
	var retlist []message.BasicMessage
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
