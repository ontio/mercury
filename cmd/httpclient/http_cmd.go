package httpclient

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/packager"
	"git.ont.io/ontid/otf/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"git.ont.io/ontid/otf/cmd"
	"github.com/urfave/cli"
)

var InvitationCmd = cli.Command{
	Name:        "invitation",
	Usage:       "new invitation",
	Description: "Generate Invitation",
	Action:      NewInvitation,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.InvitationFlag,
	},
}

var ConnectCmd = cli.Command{
	Name:        "connect",
	Usage:       "connect",
	Description: "Connect  Data",
	Action:      Connection,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.ConnectionFlag,
	},
}

var SendMsgCmd = cli.Command{
	Name:        "sendmsg",
	Usage:       "send msg ",
	Description: "Send Msg  data",
	Action:      SendMsg,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendMsgFlag,
	},
}

var ReqCredentialCmd = cli.Command{
	Name:   "reqCredential",
	Usage:  "req Credential",
	Action: ReqCredential,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendCredentialCmd,
	},
}

var ReqPresentationCmd = cli.Command{
	Name:   "reqPresentation",
	Usage:  "req presentation data",
	Action: ReqPresentation,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendCredentialCmd,
	},
}
var QueryCredCmd = cli.Command{
	Name:                   "querycredential",
	Usage:                  "query a stored credential",
	Description:            "query a stored credential",
	Action:                 QueryCredential,
	Flags:                  []cli.Flag{
		cmd.DidFlag,
		cmd.CredentialIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
	},

}
var QueryPresentationCmd = cli.Command{
	Name:                   "querypresentation",
	Usage:                  "query a stored presentation",
	Description:            "query a stored presentation",
	Action:                 QueryPresentation,
	Flags:                  []cli.Flag{
		cmd.DidFlag,
		cmd.PresentationIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.ToDID,
	},

}
var ontsdk *ontology_go_sdk.OntologySdk
var defaultAcct *ontology_go_sdk.Account
func initsdk(addr string)  {
	ontsdk = ontology_go_sdk.NewOntologySdk()
	ontsdk.NewRpcClient().SetAddress(addr)
	var err error
	defaultAcct, err = utils.OpenAccount(cmd.DEFAULT_WALLET_PATH, ontsdk)
	if err != nil {
		panic(err)
	}
}

func NewInvitation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.InvitationFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	body, err := HttpPostData(url, data)
	if err != nil {
		return fmt.Errorf("NewInvitation err:%s", err)
	}
	msg := &message.Invitation{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return fmt.Errorf("NewInvatation unmarshal err:%s", err)
	}
	fmt.Printf("msg:%v\n", msg)
	return nil
}

func Connection(ctx *cli.Context) error {
	return nil
}

func SendMsg(ctx *cli.Context) error {
	return nil
}

func ReqCredential(ctx *cli.Context) error {
	return nil
}

func ReqPresentation(ctx *cli.Context) error {
	return nil
}

func QueryCredential(ctx *cli.Context)error{

	did := ctx.String(cmd.GetFlagName(cmd.DidFlag))
	id := ctx.String(cmd.GetFlagName(cmd.CredentialIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	initsdk(url)
	req := message.QueryCredentialRequest{
		DId: did,
		Id:  id,
	}
	reqdata,err := json.Marshal(req)
	if err != nil {
		return err
	}
	env := packager.Envelope{}
	env.Message = &packager.MessageData{
		Data: reqdata,
		Sign: nil,
	}
	return nil
}

func QueryPresentation(ctx cli.Context)error {
	return nil
}
func HttpPostData(url, data string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,
			DisableKeepAlives:     false,
			IdleConnTimeout:       time.Second * 300,
			ResponseHeaderTimeout: time.Second * 300,
		},
		Timeout: time.Second * 300,
	}
	resp, err := client.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read  response body err:%s", err)
	}
	return body, nil
}
