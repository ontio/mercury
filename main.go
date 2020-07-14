package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"git.ont.io/ontid/otf/cmd"
	http_cmd "git.ont.io/ontid/otf/cmd/httpclient"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/rest"
	"git.ont.io/ontid/otf/service"
	store "git.ont.io/ontid/otf/store/leveldb"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri/ontdid"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "agent otf"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.HttpIpFlag,
		cmd.HttpPortFlag,
		cmd.ChainAddrFlag,
		cmd.HttpsPortFlag,
		cmd.EnableHttpsFlag,
		cmd.EnablePackageFlag,
	}
	app.Commands = []cli.Command{
		cmd.DidCommand,
		cmd.AddServiceCommand,
		cmd.QueryDidDocCommand,
		cmd.QueryServiceEndPointCommand,
		http_cmd.InvitationCmd,
		http_cmd.ConnectCmd,
		http_cmd.SendMsgCmd,
		http_cmd.ReqCredentialCmd,
		http_cmd.ReqPresentationCmd,
		http_cmd.QueryCredCmd,
		http_cmd.QueryPresentationCmd,
		http_cmd.QueryMsgCmd,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}
func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func startAgent(ctx *cli.Context) {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.ChainAddrFlag)))
	account, err := utils.OpenAccount(cmd.DEFAULT_WALLET_PATH, ontSdk)
	if err != nil {
		panic(err)
	}
	var port string
	if ctx.Bool(cmd.GetFlagName(cmd.EnableHttpsFlag)) {
		port = ctx.String(cmd.GetFlagName(cmd.HttpsPortFlag))
	} else {
		port = ctx.String(cmd.GetFlagName(cmd.HttpPortFlag))
	}
	if ctx.Bool(cmd.GetFlagName(cmd.EnablePackageFlag)) {
		rest.EnablePackage = true
	}
	ip := ctx.String(cmd.GetFlagName(cmd.HttpIpFlag))
	prov := store.NewProvider(cmd.DEFAULT_STORE_DIR)
	db, err := prov.OpenStore(cmd.DEFAULT_STORE_DIR)
	if err != nil {
		panic(err)
	}
	cfg := &config.Cfg{
		Port: port,
		Ip:   ip,
	}
	r := rest.InitRouter()
	ontvdri := ontdid.NewOntVDRI(ontSdk, account, "")
	msgSvr := service.NewMessageService(ontvdri, ontSdk, account, ctx.Bool(cmd.GetFlagName(cmd.EnablePackageFlag)))
	rest.NewService(account, cfg, db, msgSvr, ontvdri, ontSdk)
	middleware.Log.Infof("start agent svr%s,port:%s", account.Address, cfg.Port)
	startPort := ip + ":" + port
	if ctx.Bool(cmd.GetFlagName(cmd.EnableHttpsFlag)) {
		err = r.RunTLS(startPort, cmd.DEFAULT_CERT_PATH, cmd.DEFAULT_KEY_PATH)
		if err != nil {
			panic(err)
		}
	} else {
		err = r.Run(startPort)
		if err != nil {
			panic(err)
		}
	}
	signalHandle()
}

func signalHandle() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("get a signal: stop the rest gateway process", si.String())
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
