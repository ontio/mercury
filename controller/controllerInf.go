package controller

type ParameterInf interface {
	GetParameter() interface{}
}

type ControllerResp interface {
	GetString() (string, error)
	GetBytes() ([]byte, error)
	GetInt64() (int64, error)
	GetMap() (map[string]interface{}, error)
	GetNextMessage() (Message, error)
}

type ControllerInf interface {
	Name() string
	Initiate(param ParameterInf) error
	Process(msg Message) (ControllerResp, error)
	Shutdown() error
}
