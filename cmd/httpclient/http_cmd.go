package httpclient

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"git.ont.io/ontid/otf/cmd"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/packager"
	"git.ont.io/ontid/otf/packager/ecdsa"
	"git.ont.io/ontid/otf/utils"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
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
		cmd.FromDID,
		cmd.ToDID,
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
		cmd.FromDID,
		cmd.ToDID,
	},
}

var ReqCredentialCmd = cli.Command{
	Name:   "reqcredential",
	Usage:  "req Credential",
	Action: ReqCredential,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendCredentialCmd,
		cmd.FromDID,
		cmd.ToDID,
	},
}

var ReqPresentationCmd = cli.Command{
	Name:   "reqpresentation",
	Usage:  "req presentation data",
	Action: ReqPresentation,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendCredentialCmd,
		cmd.FromDID,
		cmd.ToDID,
	},
}
var QueryCredCmd = cli.Command{
	Name:                   "querycredential",
	Usage:                  "query a stored credential",
	Description:            "query a stored credential",
	Action:                 QueryCredential,
	Flags:                  []cli.Flag{
		cmd.CredentialIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
		cmd.ToDID,
	},

}
var QueryPresentationCmd = cli.Command{
	Name:                   "querypresentation",
	Usage:                  "query a stored presentation",
	Description:            "query a stored presentation",
	Action:                 QueryPresentation,
	Flags:                  []cli.Flag{
		cmd.PresentationIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
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
func initPackager(addr string) *ecdsa.Packager {
	ontsdk = ontology_go_sdk.NewOntologySdk()
	ontsdk.NewRpcClient().SetAddress(addr)
	acc, err := utils.OpenAccount(cmd.DEFAULT_WALLET_PATH, ontsdk)
	if err != nil {
		panic(err)
	}
	return ecdsa.New(ontsdk, acc)
}

func NewInvitation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.InvitationFlag))
	body, err := HttpPostData(ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), data)
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
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	pack := initPackager(restUrl)
	dataMsg, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("DecodeString err:%s", err)
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data: dataMsg,
		},
		MsgType: int(message.ConnectionRequestType),
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	_, err = HttpPostData(ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), hex.EncodeToString(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	return nil
}

func SendMsg(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	pack := initPackager(restUrl)
	dataMsg, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("DecodeString err:%s", err)
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data: dataMsg,
		},
		MsgType: int(message.SendGeneralMsgType),
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	_, err = HttpPostData(ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), hex.EncodeToString(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	return nil
}

func ReqCredential(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	pack := initPackager(restUrl)
	dataMsg, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("DecodeString err:%s", err)
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data: dataMsg,
		},
		MsgType: int(message.RequestCredentialType),
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	_, err = HttpPostData(ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), hex.EncodeToString(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	return nil
}

func ReqPresentation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	pack := initPackager(restUrl)
	dataMsg, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("DecodeString err:%s", err)
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data: dataMsg,
		},
		MsgType: int(message.RequestPresentationType),
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	_, err = HttpPostData(ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), hex.EncodeToString(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	return nil
	return nil
}

func QueryCredential(ctx *cli.Context)error{

	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	id := ctx.String(cmd.GetFlagName(cmd.CredentialIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	initsdk(url)
	req := message.QueryCredentialRequest{
		DId: fromdid,
		Id:  id,
	}
	reqdata,err := json.Marshal(req)
	if err != nil {
		return err
	}
	env := &packager.Envelope{}
	env.Message = &packager.MessageData{
		Data: reqdata,
		Sign: nil,
	}
	env.FromDID = fromdid
	env.ToDID = todid
	env.MsgType = int(message.QueryCredentialType)

	packer := ecdsa.New(ontsdk,defaultAcct)
	data,err := packer.PackMessage(env)
	if err != nil {
		return err
	}
	respbts,err := HttpPostData(url,string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============credential==============")
	fmt.Printf("%s\n",respbts)
	fmt.Println("==============credential==============")

	return nil
}

func QueryPresentation(ctx cli.Context)error {
	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	id := ctx.String(cmd.GetFlagName(cmd.PresentationIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	initsdk(url)


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
