package service

import (
	"fmt"
	"git.ont.io/ontid/otf/message"
)

type Customcontroller struct {
}

func NewCustomcontroller() Customcontroller {
	s := Customcontroller{}
	s.Initiate(nil)
	return s
}

func (s Customcontroller) Name() string {
	return "customcontroller"
}

func (s Customcontroller) Initiate(param ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s Customcontroller) Process(msg message.Message) (ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.InvitationType,
		message.ConnectionRequestType,
		message.ConnectionResponseType,
		message.ConnectionACKType:
		return skipmessage(msg)

	case message.ProposalCredentialType:
	case message.OfferCredentialType:
	case message.RequestCredentialType:
	case message.IssueCredentialType:
	case message.CredentialACKType:

	case message.RequestPresentationType:
	case message.PresentationType:
	case message.PresentationACKType:

	default:
		resp := ServiceResp{}
		return resp, nil
	}

	return nil, nil
}
func (s Customcontroller) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}
