package httpclient

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
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
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.ConnectionFlag,
	},
}

var SendMsgCmd = cli.Command{
	Name:        "sendmg",
	Usage:       "send msg ",
	Description: "Send Msg  data",
	Action:      SendMsg,
	Flags: []cli.Flag{
		cmd.HttpClientFlag,
		cmd.WalletFileFlag,
		cmd.SendMsgFlag,
	},
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
