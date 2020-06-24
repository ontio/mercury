package service

import "git.ont.io/ontid/otf/message"

type ParameterInf interface {
	GetParameter() interface{}
}

type ControllerResp interface {
	GetString() (string, error)
	GetBytes() ([]byte, error)
	GetInt64() (int64, error)
	GetMap() (map[string]interface{}, error)
	GetNextMessage() (message.Message, error)
	GetOriginMessage() (message.Message, error)
	GetJsonbytes()([]byte,error)
}

type ControllerInf interface {
	Name() string
	Initiate(param ParameterInf) error
	Process(msg message.Message) (ControllerResp, error)
	Shutdown() error
}
