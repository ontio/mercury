package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/rest"
	"git.ont.io/ontid/otf/utils"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "agent otf"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		utils.LogLevelFlag,
		utils.LogDirFlag,
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
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.Use(gin.Recovery())
	account, err := utils.OpenAccount(utils.DEFAULT_WALLET_PATH)
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/invitation", rest.Invite)
		v1.POST("/connection", rest.Connect)
		v1.POST("/connectionack",rest.ConnectAck)
		v1.POST("/proposalcredential",rest.ProposalCredentialReq)
		v1.POST("/sendcredential", rest.SendCredential)
		v1.POST("/issuecredentail", rest.IssueCredential)
		v1.POST("/credentialack",rest.CredentialAckInfo)
		v1.POST("/requestproof", rest.RequestProof)
		v1.POST("/presentproof", rest.PresentProof)
		v1.POST("/presentationack",rest.PresentationACKInfo)
	}
	rest.NewService()
	middleware.Log.Infof("start agent svr%s", account.Address)
	err = r.Run(utils.DEFAULT_HTTP_PORT)
	if err != nil {
		panic(err)
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
