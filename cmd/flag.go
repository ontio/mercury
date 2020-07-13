package cmd

import (
	"github.com/urfave/cli"
	"strings"
)

const (
	DEFAULT_WALLET_PATH       = "./wallet.dat"
	DEFAULT_LOG_LEVEL         = 1
	DEFAULT_HTTP_PORT         = "8080"
	DEFAULT_HTTPS_PORT        = "8443"
	DEFAULT_HTTP_IP           = "127.0.0.1"
	DEFAULT_LOG_FILE_PATH     = "./Log/"
	DEFAULT_STORE_DIR         = "./db_otf/"
	DEFAULT_BLOCK_CHAIN_ADDR  = "http://polaris2.ont.io:20336"
	DEFAULT_RPC_URL           = "http://polaris2.ont.io:20336"
	MIN_TRANSACTION_GAS       = 20000
	DEFAULT_GAS_PRICE         = 2500
	DEFAULT_WALLET_FILE_NAME  = "./wallet.dat"
	DEFAULT_DID               = ""
	DEFAULT_SERVICE_ID        = ""
	DEFAULT_TYPE              = ""
	DEFAULT_SERVICE_END_POINT = ""
	DEFAULT_INDEX             = 1
	DEFAULT_CERT_PATH         = "./key/ssl.crt"
	DEFAULT_KEY_PATH          = "./key/ssl.key"
	DEFAULT_CONNECT_DATA      = ""
	DEFAULT_SEND_MSG_DATA     = ""
	DEFAULT_CLIENT_REST_URL   = "http://127.0.0.1:8080"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: DEFAULT_LOG_LEVEL,
	}
	HttpPortFlag = cli.StringFlag{
		Name:  "http-port",
		Usage: "Set http rest port default:8080",
		Value: DEFAULT_HTTP_PORT,
	}
	HttpsPortFlag = cli.StringFlag{
		Name:  "https-port",
		Usage: "Set https rest port default:8443",
		Value: DEFAULT_HTTPS_PORT,
	}
	EnableHttpsFlag = cli.BoolFlag{
		Name:  "enable-https",
		Usage: "start https restful service",
	}
	HttpIpFlag = cli.StringFlag{
		Name:  "rest-ip",
		Usage: "Set http rest ip addr default:127.0.0.1",
		Value: DEFAULT_HTTP_IP,
	}
	ChainAddrFlag = cli.StringFlag{
		Name:  "chain-addr",
		Usage: "Set block chain rpc addr default:127.0.0.1:20334",
		Value: DEFAULT_BLOCK_CHAIN_ADDR,
	}
	RPCPortFlag = cli.StringFlag{
		Name:  "http rpc url",
		Usage: "Json rpc server listening port `<url>`",
		Value: DEFAULT_RPC_URL,
	}
	TransactionGasPriceFlag = cli.Uint64Flag{
		Name:  "gasprice",
		Usage: "Gas price of transaction",
		Value: DEFAULT_GAS_PRICE,
	}
	TransactionGasLimitFlag = cli.Uint64Flag{
		Name:  "gaslimit",
		Usage: "Gas limit of the transaction",
		Value: MIN_TRANSACTION_GAS,
	}
	WalletFileFlag = cli.StringFlag{
		Name:  "wallet,w",
		Value: DEFAULT_WALLET_FILE_NAME,
		Usage: "Wallet `<file>`",
	}
	DidFlag = cli.StringFlag{
		Name:  "did",
		Usage: "did value",
		Value: DEFAULT_DID,
	}
	ServiceIdFlag = cli.StringFlag{
		Name:  "service_id",
		Usage: "service id",
		Value: DEFAULT_SERVICE_ID,
	}
	TypeFlag = cli.StringFlag{
		Name:  "type",
		Usage: "type value",
		Value: DEFAULT_TYPE,
	}
	ServiceEndPointFlag = cli.StringFlag{
		Name:  "endpoint",
		Usage: "service end point",
		Value: DEFAULT_SERVICE_END_POINT,
	}
	IndexFlag = cli.Uint64Flag{
		Name:  "index",
		Usage: "index number",
		Value: DEFAULT_INDEX,
	}
	EnablePackageFlag = cli.BoolFlag{
		Name:  "enable-package",
		Usage: "start package msg",
	}
	HttpClientFlag = cli.StringFlag{
		Name:  "restful",
		Usage: "set http client restful url default:http://127.0.0.1:8080",
		Value: DEFAULT_CLIENT_REST_URL,
	}
	InvitationFlag = cli.StringFlag{
		Name:  "invitation data",
		Usage: "invitation data",
	}
	ConnectionFlag = cli.StringFlag{
		Name:  "connect  data",
		Usage: "connect data",
		Value: DEFAULT_CONNECT_DATA,
	}
	SendMsgFlag = cli.StringFlag{
		Name:  "send msg",
		Usage: "send msg data",
		Value: DEFAULT_SEND_MSG_DATA,
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
