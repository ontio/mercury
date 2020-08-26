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

package cmd

import (
	"github.com/urfave/cli"
	"strings"
)

const (
	DEFAULT_WALLET_PATH           = "./wallet.dat"
	DEFAULT_LOG_LEVEL             = 1
	DEFAULT_HTTP_PORT             = "8080"
	DEFAULT_HTTPS_PORT            = "8443"
	DEFAULT_HTTP_IP               = "127.0.0.1"
	DEFAULT_LOG_FILE_PATH         = "./Log/"
	DEFAULT_STORE_DIR             = "./db_mercury/"
	DEFAULT_BLOCK_CHAIN_REST_URL  = "http://polaris2.ont.io:20334"
	DEFAULT_BLOCK_CHAIN_RPC_URL   = "http://polaris2.ont.io:20336"
	MIN_TRANSACTION_GAS           = 20000
	DEFAULT_GAS_PRICE             = 2500
	DEFAULT_WALLET_FILE_NAME      = "./wallet.dat"
	DEFAULT_DID                   = ""
	DEFAULT_SERVICE_ID            = ""
	DEFAULT_TYPE                  = ""
	DEFAULT_SERVICE_END_POINT     = ""
	DEFAULT_INDEX                 = 1
	DEFAULT_CERT_PATH             = "./common/key/ssl.crt"
	DEFAULT_KEY_PATH              = "./common/key/ssl.key"
	DEFAULT_CONNECT_DATA          = ""
	DEFAULT_SEND_MSG_DATA         = ""
	DEFAULT_CLIENT_REST_URL       = "http://127.0.0.1:8080"
	DEFAULT_REQ_CREDENTIAL_DATA   = ""
	DEFAULT_REQ_PRESENTATION_DATA = ""
	DEFAULT_CHECK_LOG             = 6
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: DEFAULT_LOG_LEVEL,
	}
	LogDirFlag = cli.StringFlag{
		Name:  "log-dir",
		Usage: "log output to the file",
		Value: DEFAULT_LOG_FILE_PATH,
	}
	DisableLogFileFlag = cli.BoolFlag{
		Name:  "disable-log-file",
		Usage: "Discard log output to file",
	}
	HttpPortFlag = cli.StringFlag{
		Name:  "http-port",
		Usage: "Set http rest port",
		Value: DEFAULT_HTTP_PORT,
	}
	HttpsPortFlag = cli.StringFlag{
		Name:  "https-port",
		Usage: "set https rest port",
		Value: DEFAULT_HTTPS_PORT,
	}
	SelfDIDFlag = cli.StringFlag{
		Name:  "agent-did",
		Usage: "agent did",
		//Required: true,
	}
	EnableHttpsFlag = cli.BoolFlag{
		Name:  "enable-https",
		Usage: "start https restful service",
	}
	HttpIpFlag = cli.StringFlag{
		Name:  "rest-ip",
		Usage: "set services http rest ip addr",
		Value: DEFAULT_HTTP_IP,
	}
	RestUrlFlag = cli.StringFlag{
		Name:  "rest-url",
		Usage: "set block chain rest url",
		Value: DEFAULT_BLOCK_CHAIN_REST_URL,
	}
	RpcUrlFlag = cli.StringFlag{
		Name:  "chain-rpc-url",
		Usage: "Set block chain rpc url",
		Value: DEFAULT_BLOCK_CHAIN_RPC_URL,
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
		Name:  "service-id",
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
		Usage: "set http client restful url",
		Value: DEFAULT_CLIENT_REST_URL,
	}
	InvitationFlag = cli.StringFlag{
		Name:  "invitation-data",
		Usage: "invitation data",
	}
	ConnectionFlag = cli.StringFlag{
		Name:  "connect-data",
		Usage: "connect data",
		Value: DEFAULT_CONNECT_DATA,
	}
	SendMsgFlag = cli.StringFlag{
		Name:  "send-msg",
		Usage: "send msg data",
		Value: DEFAULT_SEND_MSG_DATA,
	}
	ReadLatestMsgFlag = cli.BoolFlag{
		Name:  "latest",
		Usage: "read latest message",
	}
	RemoveAfterReadFlag = cli.BoolFlag{
		Name:  "remove-after-read",
		Usage: "remove after read",
	}
	ReqCredentialCmd = cli.StringFlag{
		Name:  "req-credential",
		Usage: "req credential data",
		Value: DEFAULT_REQ_CREDENTIAL_DATA,
	}
	ReqPresentationCmd = cli.StringFlag{
		Name:  "req-presentation",
		Usage: "req presentation data",
		Value: DEFAULT_REQ_PRESENTATION_DATA,
	}
	CredentialIdFlag = cli.StringFlag{
		Name:  "credential-id",
		Usage: "credential id",
	}
	PresentationIdFlag = cli.StringFlag{
		Name:  "presentation-id",
		Usage: "presentation id",
	}
	FromDID = cli.StringFlag{
		Name:  "from-did",
		Usage: "from did",
	}
	ToDID = cli.StringFlag{
		Name:  "to-did",
		Usage: "to did",
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
