package httpclient

import (
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
	sdk "github.com/ontio/ontology-go-sdk"
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
		cmd.FromDID,
		cmd.ToDID,
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
		cmd.FromDID,
		cmd.ToDID,
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
		cmd.FromDID,
		cmd.ToDID,
		cmd.SendMsgFlag,
	},
}

var QueryMsgCmd = cli.Command{
	Name:        "querymsg",
	Description: "query general message",
	Action:      QueryMsg,
	Flags: []cli.Flag{
		cmd.RPCPortFlag,
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.FromDID,
		cmd.ToDID,
		cmd.ReadLatestMsgFlag,
		cmd.RemoveAfterReadFlag,
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
		cmd.FromDID,
		cmd.ToDID,
		cmd.ReqCredentialCmd,
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
		cmd.FromDID,
		cmd.ToDID,
		cmd.ReqPresentationCmd,
	},
}
var QueryCredCmd = cli.Command{
	Name:        "querycredential",
	Usage:       "query a stored credential",
	Description: "query a stored credential",
	Action:      QueryCredential,
	Flags: []cli.Flag{
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
		cmd.ToDID,
		cmd.CredentialIdFlag,
	},
}
var QueryPresentationCmd = cli.Command{
	Name:        "querypresentation",
	Usage:       "query a stored presentation",
	Description: "query a stored presentation",
	Action:      QueryPresentation,
	Flags: []cli.Flag{
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
		cmd.ToDID,
		cmd.PresentationIdFlag,
	},
}

func initPackager(addr string) *ecdsa.Packager {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(addr)
	acc, err := utils.OpenAccount(cmd.DEFAULT_WALLET_PATH, ontSdk)
	if err != nil {
		panic(err)
	}
	return ecdsa.New(ontSdk, acc)
}

func NewInvitation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.InvitationFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.InvitationType)
	invite := &message.Invitation{}
	err := json.Unmarshal([]byte(data), invite)
	if err != nil {
		return err
	}
	reqData, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.InvitationType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	pack := initPackager(restUrl)
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	body, err := HttpPostData(url, string(bys))
	if err != nil {
		return fmt.Errorf("NewInvitation err:%s", err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func Connection(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	invite := &message.ConnectionRequest{}
	err := json.Unmarshal([]byte(data), invite)
	if err != nil {
		return err
	}
	reqData, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.ConnectionRequestType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	pack := initPackager(restUrl)
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.ConnectionRequestType)
	body, err := HttpPostData(url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func SendMsg(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.SendMsgFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	basicMsg := &message.BasicMessage{}
	err := json.Unmarshal([]byte(data), basicMsg)
	if err != nil {
		return err
	}
	reqData, err := json.Marshal(basicMsg)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.SendGeneralMsgType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	pack := initPackager(restUrl)
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.SendGeneralMsgType)
	body, err := HttpPostData(url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func QueryMsg(ctx *cli.Context) error {
	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	latest := ctx.Bool(cmd.GetFlagName(cmd.ReadLatestMsgFlag))
	rar := ctx.Bool(cmd.GetFlagName(cmd.RemoveAfterReadFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	packer := initPackager(restUrl)

	req := message.QueryGeneralMessageRequest{
		DID:             fromdid,
		Latest:          latest,
		RemoveAfterRead: rar,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return err
	}

	env := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqdata,
			MsgType: int(message.QueryGeneralMessageType),
		},
		FromDID: fromdid,
		ToDID:   todid,
	}
	data, err := packer.PackMessage(env)
	if err != nil {
		return err
	}
	url = url + utils.GetApiName(message.QueryGeneralMessageType)
	respbts, err := HttpPostData(url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============general message==============")
	fmt.Printf("%s\n", respbts)
	fmt.Println("==============general message==============")
	return nil
}

func ReqCredential(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ReqCredentialCmd))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	invite := &message.RequestCredential{}
	err := json.Unmarshal([]byte(data), invite)
	if err != nil {
		return err
	}
	reqData, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.RequestCredentialType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	pack := initPackager(restUrl)
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.RequestCredentialType)
	body, err := HttpPostData(url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func ReqPresentation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ReqPresentationCmd))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	invite := &message.RequestPresentation{}
	err := json.Unmarshal([]byte(data), invite)
	if err != nil {
		return err
	}
	reqData, err := json.Marshal(invite)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.RequestPresentationType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	pack := initPackager(restUrl)
	bys, err := pack.PackMessage(msg)
	if err != nil {
		return fmt.Errorf("packMessage err:%s", err)
	}
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.RequestPresentationType)
	body, err := HttpPostData(url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func QueryCredential(ctx *cli.Context) error {
	id := ctx.String(cmd.GetFlagName(cmd.CredentialIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	req := message.QueryCredentialRequest{
		DId: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		Id:  id,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.QueryCredentialType),
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	packer := initPackager(rpc)
	data, err := packer.PackMessage(msg)
	if err != nil {
		return err
	}
	url = url + utils.GetApiName(message.QueryCredentialType)
	body, err := HttpPostData(url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============credential==============")
	fmt.Printf("%s\n", body)
	fmt.Println("==============credential==============")
	return nil
}

func QueryPresentation(ctx *cli.Context) error {
	id := ctx.String(cmd.GetFlagName(cmd.PresentationIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))

	req := message.QueryPresentationRequest{
		DId: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		Id:  id,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	msg := &packager.Envelope{
		Message: &packager.MessageData{
			Data:    reqData,
			MsgType: int(message.QueryPresentationType),
			Sign:    nil,
		},
		FromDID: ctx.String(cmd.GetFlagName(cmd.FromDID)),
		ToDID:   ctx.String(cmd.GetFlagName(cmd.ToDID)),
	}
	packer := initPackager(rpc)
	data, err := packer.PackMessage(msg)
	if err != nil {
		return err
	}
	url = url + utils.GetApiName(message.QueryPresentationType)

	body, err := HttpPostData(url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============presentation==============")
	fmt.Printf("%s\n", body)
	fmt.Println("==============presentation==============")

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
	return ioutil.ReadAll(resp.Body)
}
