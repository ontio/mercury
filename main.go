package main

import (
	"fmt"
	"git.ont.io/ontid/otf/config"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/rest"
	"git.ont.io/ontid/otf/utils"
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
	account, err := utils.OpenAccount(utils.DEFAULT_WALLET_PATH)
	if err != nil {
		panic(err)
	}
	r := rest.InitRouter()
	cfg := &config.Cfg{
		Port: utils.DEFAULT_HTTP_PORT,
		Ip:   "",
	}
	rest.NewService(account,cfg)

	middleware.Log.Infof("start agent svr%s,port:%s", account.Address, cfg.Port)
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
