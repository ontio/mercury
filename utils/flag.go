package utils

import (
	"github.com/urfave/cli"
	"strings"
)

const (
	DEFAULT_WALLET_PATH = "./wallet.dat"
	//DEFAULT_LOG_LEVEL                       = log.InfoLo
	DEFAULT_HTTP_PORT     = ":8080"
	DEFAULT_LOG_FILE_PATH = "./Log/"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		//Value: config.DEFAULT_LOG_LEVEL,
	}
	LogDirFlag = cli.StringFlag{
		Name:  "log-dir",
		Usage: "log output to the file",
		//Value: log.PATH,
	}
)

//GetFlagName deal with short flag, and return the flag name whether flag name have short name
func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
