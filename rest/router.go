package rest

import (
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/service"
)

var (
	Svr *service.Service
)

func NewService() {
	Svr := service.NewService()
	Svr.RegisterController(service.NewSyscontroller())
	Svr.RegisterController(service.NewCustomcontroller())
}

func SendMsg(msgType message.MessageType, data map[string]interface{}) (interface{}, error) {
	msg := message.Message{MessageType: msgType, Content: data}
	resp, err := Svr.Serv(msg)
	if err != nil {
		middleware.Log.Errorf("err:%s", err)
		return nil, err
	}
	return resp, nil
}
