package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"git.ont.io/ontid/otf/utils"
)

func main() {
	r := gin.Default()
	r.Use()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
	account,err := utils.OpenAccount("./wallet.dat")
	if err != nil {
		panic(err)
	}
	fmt.Println("addr:",account.Address)
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