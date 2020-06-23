package service

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/did"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"github.com/google/uuid"
	"github.com/itchyny/base58-go"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	Version        = "1.0"
	InvitationSpec = "spec/connections/" + Version + "/invitation"
)

type Syscontroller struct {
	account *sdk.Account
	did     did.Did
	cfg     *config.Cfg
}

func NewSyscontroller(acct *sdk.Account,cfg *config.Cfg) Syscontroller {

	did := did.NewOntDID(cfg,acct)
	s := Syscontroller{
		account:acct,
		did:did,
		cfg:cfg,
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

		return ServiceResp{
			OriginalMessage: msg,
			Message:         m,
		}, nil

	case message.ConnectionRequestType:
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
	invitaion.ServiceEndpoint = "http://ip:port"

	sigdata, err := s.sign([]byte(s.did.String() + invitaion.Id))
	if err != nil {
		return nil, err
	}

	receipkey, err := base58.BitcoinEncoding.Encode(sigdata)
	if err != nil {
		return nil, err
	}

	invitaion.RecipientKeys = []string{string(receipkey)}
	invitaion.RoutingKeys = []string{string(receipkey)}

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
	jsonbytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})

	err = json.Unmarshal(jsonbytes, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
