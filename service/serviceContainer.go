package service

import (
	"container/list"
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
)

type ServiceInf interface {
	RegisterController(name string, c ControllerInf) error
	GetController(name string) (ControllerInf, error)
	GetControllers() (*list.List, error)
	RemoveController(name string) error
	Serv(message message.Message) (ControllerResp, error)
}

type ServiceResp struct {
	OriginalMessage message.Message
	Message         interface{}
	JsonBytes      []byte
}

func (r ServiceResp) GetString() (string, error) {
	j, err := json.Marshal(r.Message)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
func (r ServiceResp) GetBytes() ([]byte, error) {
	return json.Marshal(r.Message)
}

func (r ServiceResp) GetInt64() (int64, error) {
	return -1, fmt.Errorf("not support")
}

func (r ServiceResp) GetMap() (map[string]interface{}, error) {
	return nil, fmt.Errorf("not support")
}

func (r ServiceResp) GetNextMessage() (message.Message, error) {
	m := message.Message{}
	m.MessageType = r.OriginalMessage.MessageType
	m.Content = r.Message
	m.JsonBytes = r.JsonBytes
	return m, nil
}

func (r ServiceResp) GetOriginMessage() (message.Message, error) {
	return r.OriginalMessage, nil
}

func (r ServiceResp)GetJsonbytes()([]byte,error){
	return r.JsonBytes,nil
}

type Service struct {
	//store
	Container *list.List
}

func NewService() *Service {
	return &Service{Container: list.New()}
}

func (s *Service) RegisterController(c ControllerInf) error {

	s.Container.PushBack(c)
	return nil
}

func (s *Service) GetController(name string) (ControllerInf, error) {

	for e := s.Container.Front(); e != nil; e = e.Next() {
		c := e.Value.(ControllerInf)
		if c.Name() == name {
			return c, nil
		}
	}
	return nil, nil
}

func (s *Service) GetControllers() (*list.List, error) {

	return s.Container, nil
}

func (s *Service) RemoveController(name string) error {

	for e := s.Container.Front(); e != nil; e = e.Next() {
		c := e.Value.(ControllerInf)
		if c.Name() == name {
			s.Container.Remove(e)
		}
	}

	return nil
}

func (s *Service) Serv(message message.Message) (ControllerResp, error) {

	m := message
	for e := s.Container.Front(); e != nil; e = e.Next() {
		c := e.Value.(ControllerInf)
		tmpmsg, err := c.Process(m)
		if err != nil {
			return tmpmsg, err
		}
		if e.Next() == nil {
			return tmpmsg, nil
		}

		m, err = tmpmsg.GetNextMessage()
		if err != nil {
			return tmpmsg, err
		}
	}
	//never reach here
	return ServiceResp{}, nil
}
