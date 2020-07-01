
package controller

import (
"fmt"
"git.ont.io/ontid/otf/message"
"git.ont.io/ontid/otf/service"
)

type CredentialController struct {
}

func NewCredentialController() CredentialController {
	s := CredentialController{}
	s.Initiate(nil)
	return s
}

func (s CredentialController) Name() string {
	return "CredentialController"
}

func (s CredentialController) Initiate(param service.ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s CredentialController) Process(msg message.Message) (service.ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.InvitationType,
		message.ConnectionRequestType,
		message.ConnectionResponseType,
		message.ConnectionACKType:
		return service.Skipmessage(msg)

	case message.ProposalCredentialType:
		fmt.Printf("")

	case message.OfferCredentialType:

	case message.RequestCredentialType:

	case message.IssueCredentialType:

	case message.CredentialACKType:

	case message.RequestPresentationType,
	message.PresentationType,
	message.PresentationACKType:
		service.Skipmessage(msg)

	default:
		resp := service.ServiceResp{}
		return resp, nil
	}

	return nil, nil
}
func (s CredentialController) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}
