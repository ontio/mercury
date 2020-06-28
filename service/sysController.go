package service

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/did"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/store"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	Version        = "1.0"
	InvitationSpec = "spec/connections/" + Version + "/invitation"
	ConnectionKey  = "Connection"
)

type Syscontroller struct {
	account *sdk.Account
	did     did.Did
	cfg     *config.Cfg
	store   store.Store
}

func NewSyscontroller(acct *sdk.Account, cfg *config.Cfg) Syscontroller {

	did := did.NewOntDID(cfg, acct)
	s := Syscontroller{
		account: acct,
		did:     did,
		cfg:     cfg,
	}
	s.Initiate(nil)
	return s
}

func (s Syscontroller) Name() string {
	return "syscontroller"
}

func (s Syscontroller) Initiate(param ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s Syscontroller) Process(msg message.Message) (ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	//for system
	case message.InvitationType:
		middleware.Log.Infof("resolve invitation")
		//todo verify request
		invitation ,ok:= msg.Content.(message.Invitation)
		if !ok {
			return nil,fmt.Errorf("message format is not correct")
		}
		//set uuid to invitation id
		invitation.Id = uuid.New().String()

		//store the invitation
		jsonbytes, err := json.Marshal(invitation)
		if err != nil {
			return nil, err
		}
		//save invitation
		err = s.SaveInvitation(invitation)
		if err != nil {
			return nil, err
		}

		return ServiceResp{
			OriginalMessage: msg,
			Message:         invitation,
			JsonBytes:       jsonbytes,
		}, nil
	case message.SendConnectionRequestType:
		//send connection req for agent
		middleware.Log.Infof("resolve send connection request")
		//todo verify request
		cr := msg.Content.(message.ConnectionRequest)
		cr.Id = uuid.New().String()




	case message.ConnectionRequestType:
		//middleware.Log.Infof("resolve connection request")
		//req := msg.Content.(message.ConnectionRequest)
		//req.



	case message.ConnectionResponseType:
	case message.ConnectionACKType:

	//for custom
	case message.ProposalCredentialType:
	case message.OfferCredentialType:
	case message.RequestCredentialType:
	case message.IssueCredentialType:
	case message.CredentialACKType:

	case message.RequestPresentationType:
	case message.PresentationType:
	case message.PresentationACKType:

	default:

	}

	resp := ServiceResp{}
	return resp, nil
}
func (s Syscontroller) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}

//func (s Syscontroller) generateInvitation() (*message.Invitation, error) {
//	invitaion := new(message.Invitation)
//	invitaion.Type = fmt.Sprintf("%s;%s", s.did.String(), InvitationSpec)
//	invitaion.Id = uuid.New().String()
//	//fixme to set a lable
//	invitaion.Label = s.account.Address.ToBase58()
//	invitaion.Did = s.did.String()
//	addrbase58 := s.account.Address.ToBase58()
//
//
//	return invitaion, nil
//}

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

//

func (s Syscontroller) SaveInvitation(iv message.Invitation) error {
	key := fmt.Sprintf("%s_%s", ConnectionKey, iv.Id)
	tmp,err :=s.store.Get(key)
	if err != nil{
		return err
	}
	if tmp != nil{
		return fmt.Errorf("invitation with id:%s existed",iv.Id)
	}

	rec := InvitationRec{
		Invitation: iv,
		State:      ConnectionInit,
	}

	bs ,err:= json.Marshal(rec)
	if err != nil{
		return err
	}

	return s.store.Put(key,bs)

}
