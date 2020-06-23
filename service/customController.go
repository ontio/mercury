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
	case message.Invitation:
	case message.ConnectionRequest:
	case message.ConnectionResponse:
	case message.ConnectionACK:

	case message.ProposalCredential:
	case message.OfferCredential:
	case message.RequestCredential:
	case message.IssueCredential:
	case message.CredentialACK:

	case message.RequestPresentation:
	case message.Presentation:
	case message.PresentationACK:

	default:

	}

	resp := ServiceResp{}
	return resp, nil
}
func (s Customcontroller) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}
