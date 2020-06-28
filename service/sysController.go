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
		invitation, err := s.generateInvitation()
		if err != nil {
			return nil, err
		}

		m, err := s.toMap(invitation)
		if err != nil {
			return nil, err
		}
		jsonbytes, err := json.Marshal(invitation)
		if err != nil {
			return nil, err
		}
		return ServiceResp{
			OriginalMessage: msg,
			Message:         m,
			JsonBytes:       jsonbytes,
		}, nil

	case message.ConnectionRequestType:
		middleware.Log.Infof("resolve connection request")
		req := msg.Content.(message.ConnectionRequest)

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

func (s Syscontroller) generateInvitation() (*message.Invitation, error) {
	invitaion := new(message.Invitation)
	invitaion.Type = fmt.Sprintf("%s;%s", s.did.String(), InvitationSpec)
	invitaion.Id = uuid.New().String()
	//fixme to set a lable
	invitaion.Label = s.account.Address.ToBase58()
	invitaion.ServiceEndpoint = fmt.Sprintf("http://%s:%s", s.cfg.Ip, s.cfg.Port)
	invitaion.Did = s.did.String()
	addrbase58 := s.account.Address.ToBase58()

	invitaion.RecipientKeys = []string{addrbase58}
	invitaion.RoutingKeys = []string{addrbase58}

	return invitaion, nil
}

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

func (s Syscontroller) SaveConnectionState(id string, state ConnectionState) error {
	key := fmt.Sprintf("%s_%s", ConnectionKey, id)

	switch state {
	case ConnectionInit:
	}

	if v, _ := s.store.Get(key); v != nil {
		return fmt.Errorf("")
	}
}
