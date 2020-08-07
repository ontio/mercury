/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/ontio/mercury/cmd"
	"github.com/ontio/mercury/cmd/did"
	http_cmd "github.com/ontio/mercury/cmd/httpclient"
	"github.com/ontio/mercury/common/config"
	"github.com/ontio/mercury/common/log"
	"github.com/ontio/mercury/common/packager/ecdsa"
	"github.com/ontio/mercury/service"
	"github.com/ontio/mercury/service/common"
	store "github.com/ontio/mercury/store/leveldb"
	"github.com/ontio/mercury/utils"
	"github.com/ontio/mercury/vdri/ontdid"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "agent otf"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.LogDirFlag,
		cmd.HttpIpFlag,
		cmd.HttpPortFlag,
		cmd.RestUrlFlag,
		cmd.RpcUrlFlag,
		cmd.HttpsPortFlag,
		cmd.SelfDIDFlag,
		cmd.EnableHttpsFlag,
		cmd.EnablePackageFlag,
	}
	app.Commands = []cli.Command{
		did.DidCommand,
		http_cmd.HttpClientCmd,
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
	initLog(ctx)
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
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
		common.EnablePackage = true
	}
	selfDid := ctx.String(cmd.GetFlagName(cmd.SelfDIDFlag))
	ip := ctx.String(cmd.GetFlagName(cmd.HttpIpFlag))
	prov := store.NewProvider(cmd.DEFAULT_STORE_DIR)
	db, err := prov.OpenStore(cmd.DEFAULT_STORE_DIR)
	if err != nil {
		panic(err)
	}
	cfg := &config.Cfg{
		Port:    port,
		Ip:      ip,
		SelfDID: selfDid,
	}
	ontVdri := ontdid.NewOntVDRI(ontSdk, account, selfDid)
	msgSvr := common.NewMessageService(ontVdri, ontSdk, account, ctx.Bool(cmd.GetFlagName(cmd.EnablePackageFlag)), cfg)
	r := service.NewApiRouter(ecdsa.New(ontSdk, account), db, msgSvr, ontVdri)
	log.Infof("start agent svr account:%s,port:%s", account.Address.ToBase58(), cfg.Port)
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

func initLog(ctx *cli.Context) {
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	disableLogFile := ctx.GlobalBool(cmd.GetFlagName(cmd.DisableLogFileFlag))
	if disableLogFile {
		log.InitLog(logLevel, log.Stdout)
	} else {
		logFileDir := ctx.String(cmd.GetFlagName(cmd.LogDirFlag))
		logFileDir = filepath.Join(logFileDir, "") + string(os.PathSeparator)
		log.InitLog(logLevel, logFileDir, log.Stdout)
	}
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
