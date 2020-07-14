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
		cmd.InvitationFlag,
		cmd.FromDID,
		cmd.ToDID,
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
	Name:        "querycredential",
	Usage:       "query a stored credential",
	Description: "query a stored credential",
	Action:      QueryCredential,
	Flags: []cli.Flag{
		cmd.CredentialIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
		cmd.ToDID,
	},
}
var QueryPresentationCmd = cli.Command{
	Name:        "querypresentation",
	Usage:       "query a stored presentation",
	Description: "query a stored presentation",
	Action:      QueryPresentation,
	Flags: []cli.Flag{
		cmd.PresentationIdFlag,
		cmd.HttpClientFlag,
		cmd.RPCPortFlag,
		cmd.FromDID,
		cmd.ToDID,
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
	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	rpc := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)) + utils.GetApiName(message.InvitationType)
	env := &packager.Envelope{}
	env.Message = &packager.MessageData{
		Data:    []byte(data),
		MsgType: int(message.QueryCredentialType),
		Sign:    nil,
	}
	env.FromDID = fromdid
	env.ToDID = todid

	packer := initPackager(rpc)
	msg, err := packer.PackMessage(env)
	if err != nil {
		return err
	}
	body, err := HttpPostData(url, string(msg))
	if err != nil {
		return fmt.Errorf("NewInvitation err:%s", err)
	}
	fmt.Printf(":%s\n", body)
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
			Data:    dataMsg,
			MsgType: int(message.ConnectionRequestType),
		},
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
			Data:    dataMsg,
			MsgType: int(message.SendGeneralMsgType),
		},
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
			Data:    dataMsg,
			MsgType: int(message.RequestCredentialType),
		},
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
			Data:    dataMsg,
			MsgType: int(message.RequestPresentationType),
		},
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

func QueryCredential(ctx *cli.Context) error {

	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	id := ctx.String(cmd.GetFlagName(cmd.CredentialIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))
	req := message.QueryCredentialRequest{
		DId: fromdid,
		Id:  id,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return err
	}
	env := &packager.Envelope{}
	env.Message = &packager.MessageData{
		Data:    reqdata,
		MsgType: int(message.QueryCredentialType),
		Sign:    nil,
	}
	env.FromDID = fromdid
	env.ToDID = todid

	packer := initPackager(rpc)
	data, err := packer.PackMessage(env)
	if err != nil {
		return err
	}
	url = url + utils.GetApiName(message.QueryCredentialType)
	respbts, err := HttpPostData(url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============credential==============")
	fmt.Printf("%s\n", respbts)
	fmt.Println("==============credential==============")

	return nil
}

func QueryPresentation(ctx cli.Context) error {
	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	id := ctx.String(cmd.GetFlagName(cmd.PresentationIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RPCPortFlag))

	req := message.QueryPresentationRequest{
		DId: fromdid,
		Id:  id,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return err
	}
	env := &packager.Envelope{}
	env.Message = &packager.MessageData{
		Data:    reqdata,
		MsgType: int(message.QueryPresentationType),
		Sign:    nil,
	}
	env.FromDID = fromdid
	env.ToDID = todid

	packer := initPackager(rpc)
	data, err := packer.PackMessage(env)
	if err != nil {
		return err
	}
	respbts, err := HttpPostData(url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============presentation==============")
	fmt.Printf("%s\n", respbts)
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
	method := "POST"
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("newRequest err:%s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
