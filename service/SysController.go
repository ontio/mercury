package service

import (
	"fmt"
	"git.ont.io/ontid/otf/controller"
	"git.ont.io/ontid/otf/message"
)

type Syscontroller struct {

}

func(s *Syscontroller)Name()string{
	return "syscontroller"
}

func (s *Syscontroller)Initiate(param controller.ParameterInf)error {
	fmt.Println("%s Initiate",s.Name())
	//todo add logic
	return nil
}

func (s *Syscontroller)Process(msg message.Message) (controller.ControllerResp, error){
	fmt.Printf("%s Process:%v",s.Name(),msg)
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
	return resp,nil
}