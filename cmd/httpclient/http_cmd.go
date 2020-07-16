package httpclient

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/cmd"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/packager"
	"git.ont.io/ontid/otf/packager/ecdsa"
	"git.ont.io/ontid/otf/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

var HttpClientCmd = cli.Command{
	Action:    cli.ShowSubcommandHelp,
	Name:      "httpclient",
	Usage:     "http client cli",
	ArgsUsage: "[arguments ...]",
	Description: "cli management commands can be use to invitation,connect,sendmsg,reqcredential," +
		"reqpresentation,querycredential,querypresentation.you can use ./agent-otf httpclient --help to view information",
	Subcommands: []cli.Command{
		{
			Action:      newInvitation,
			Name:        "invitation",
			Usage:       "new invitation",
			Description: "Generate Invitation",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.InvitationFlag,
			},
		},
		{
			Action:      connection,
			Name:        "connect",
			Usage:       "connect",
			Description: "Connect  Data",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.ConnectionFlag,
			},
		},
		{
			Action:      sendMsg,
			Name:        "sendmsg",
			Usage:       "send basic msg",
			Description: "send basic msg data",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.SendMsgFlag,
			},
		},
		{
			Action:      queryMsg,
			Name:        "querymsg",
			Usage:       "query basic msg",
			Description: "query basic message",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.ReadLatestMsgFlag,
				cmd.RemoveAfterReadFlag,
			},
		},
		{
			Action:      reqCredential,
			Name:        "reqcredential",
			Usage:       "req credential",
			Description: "req credential",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.ReqCredentialCmd,
			},
		},
		{
			Action:      reqPresentation,
			Name:        "reqpresentation",
			Usage:       "req presentation data",
			Description: "req presentation data",
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.HttpClientFlag,
				cmd.WalletFileFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.ReqPresentationCmd,
			},
		},
		{
			Action:      queryCredential,
			Name:        "querycredential",
			Usage:       "query a stored credential",
			Description: "query a stored credential",
			Flags: []cli.Flag{
				cmd.HttpClientFlag,
				cmd.RpcUrlFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.CredentialIdFlag,
			},
		},
		{
			Action:      queryPresentation,
			Name:        "querypresentation",
			Usage:       "query a stored presentation",
			Description: "query a stored presentation",
			Flags: []cli.Flag{
				cmd.HttpClientFlag,
				cmd.RpcUrlFlag,
				cmd.FromDID,
				cmd.ToDID,
				cmd.PresentationIdFlag,
			},
		},
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

func newInvitation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.InvitationFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(bys))
	if err != nil {
		return fmt.Errorf("NewInvitation err:%s", err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func connection(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ConnectionFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func sendMsg(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.SendMsgFlag))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func queryMsg(ctx *cli.Context) error {
	fromdid := ctx.String(cmd.GetFlagName(cmd.FromDID))
	todid := ctx.String(cmd.GetFlagName(cmd.ToDID))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	respbts, err := utils.HttpPostData(utils.NewClient(), url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============general message==============")
	fmt.Printf("%s\n", respbts)
	fmt.Println("==============general message==============")
	return nil
}

func reqCredential(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ReqCredentialCmd))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func reqPresentation(ctx *cli.Context) error {
	data := ctx.String(cmd.GetFlagName(cmd.ReqPresentationCmd))
	restUrl := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(bys))
	if err != nil {
		return fmt.Errorf("http post url:%s err:%s", ctx.String(cmd.GetFlagName(cmd.HttpClientFlag)), err)
	}
	fmt.Printf(":%s\n", body)
	return nil
}

func queryCredential(ctx *cli.Context) error {
	id := ctx.String(cmd.GetFlagName(cmd.CredentialIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))
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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============credential==============")
	fmt.Printf("%s\n", body)
	fmt.Println("==============credential==============")
	return nil
}

func queryPresentation(ctx *cli.Context) error {
	id := ctx.String(cmd.GetFlagName(cmd.PresentationIdFlag))
	url := ctx.String(cmd.GetFlagName(cmd.HttpClientFlag))
	rpc := ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag))

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
	body, err := utils.HttpPostData(utils.NewClient(), url, string(data))
	if err != nil {
		return err
	}
	fmt.Println("==============presentation==============")
	fmt.Printf("%s\n", body)
	fmt.Println("==============presentation==============")

	return nil
}
