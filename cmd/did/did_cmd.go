package did

import (
	"fmt"

	"git.ont.io/ontid/otf/cmd"
	"git.ont.io/ontid/otf/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

var DidCommand = cli.Command{
	Action:    cli.ShowSubcommandHelp,
	Name:      "did",
	Usage:     "did cli",
	ArgsUsage: "[arguments ...]",
	Description: "cli management commands can be use to new did,addsvr,query diddoc" +
		"query endpoint",
	Subcommands: []cli.Command{
		{
			Name:        "newdid",
			Usage:       "new did then register to block chain",
			Description: "new did,then register to block chain",
			Action:      newDid,
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.TransactionGasPriceFlag,
				cmd.TransactionGasLimitFlag,
				cmd.WalletFileFlag,
				cmd.HttpsPortFlag,
				cmd.EnableHttpsFlag,
			},
		},
		{
			Name:        "addsvr",
			Usage:       "add service endpoint",
			Description: "Use Did add service endpoint to contract",
			Action:      addService,
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.TransactionGasPriceFlag,
				cmd.TransactionGasLimitFlag,
				cmd.WalletFileFlag,
				cmd.DidFlag,
				cmd.ServiceIdFlag,
				cmd.TypeFlag,
				cmd.ServiceEndPointFlag,
				cmd.IndexFlag,
			},
		},
		{
			Name:        "updatesvr",
			Usage:       "update service endpoint",
			Description: "Use Did add service endpoint to contract",
			Action:      updateService,
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.TransactionGasPriceFlag,
				cmd.TransactionGasLimitFlag,
				cmd.WalletFileFlag,
				cmd.DidFlag,
				cmd.ServiceIdFlag,
				cmd.TypeFlag,
				cmd.ServiceEndPointFlag,
				cmd.IndexFlag,
			},
		},
		{
			Name:        "diddoc",
			Usage:       "query did doc from block chain",
			Description: "query did doc from block chain",
			Action:      queryDidDoc,
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.DidFlag,
			},
		},
		{
			Name:        "endpoint",
			Usage:       "query service endPoint from block chain",
			Description: "query service endPoint from block chain",
			Action:      QueryEndPoint,
			Flags: []cli.Flag{
				cmd.RpcUrlFlag,
				cmd.DidFlag,
			},
		},
	},
}

func newDid(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
	gasPrice := ctx.Uint64(cmd.TransactionGasPriceFlag.Name)
	gasLimit := ctx.Uint64(cmd.TransactionGasLimitFlag.Name)
	optionFile := checkFileName(ctx)
	acc, err := utils.OpenAccount(optionFile, ontSdk)
	if err != nil {
		return fmt.Errorf("open account err:%s", err)
	}
	did, err := NewDID(ontSdk, acc, gasPrice, gasLimit)
	if err != nil {
		return fmt.Errorf("new did err:%s", err)
	}
	fmt.Printf("did:  %v\n", did)
	return nil
}

func addService(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
	gasPrice := ctx.Uint64(cmd.TransactionGasPriceFlag.Name)
	gasLimit := ctx.Uint64(cmd.TransactionGasLimitFlag.Name)
	did := ctx.String(cmd.GetFlagName(cmd.DidFlag))
	optionFile := checkFileName(ctx)
	acc, err := utils.OpenAccount(optionFile, ontSdk)
	if err != nil {
		return fmt.Errorf("open account err:%s", err)
	}
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return fmt.Errorf("ontsdk is nil")
	}
	serviceId := ctx.String(cmd.GetFlagName(cmd.ServiceIdFlag))
	type_ := ctx.String(cmd.GetFlagName(cmd.TypeFlag))
	serviceEndpoint := ctx.String(cmd.GetFlagName(cmd.ServiceEndPointFlag))
	index := ctx.Uint64(cmd.GetFlagName(cmd.IndexFlag))
	txHash, err := ontSdk.Native.OntId.AddService(gasPrice, gasLimit, acc, did, []byte(serviceId), []byte(type_), []byte(serviceEndpoint), uint32(index), acc)
	if err != nil {
		return fmt.Errorf("add service err:%s", err)
	}
	fmt.Printf("txHash:%v\n", txHash.ToHexString())
	return nil
}

func updateService(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
	gasPrice := ctx.Uint64(cmd.TransactionGasPriceFlag.Name)
	gasLimit := ctx.Uint64(cmd.TransactionGasLimitFlag.Name)
	did := ctx.String(cmd.GetFlagName(cmd.DidFlag))
	optionFile := checkFileName(ctx)
	acc, err := utils.OpenAccount(optionFile, ontSdk)
	if err != nil {
		return fmt.Errorf("open account err:%s", err)
	}
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return fmt.Errorf("ontsdk is nil")
	}
	serviceId := ctx.String(cmd.GetFlagName(cmd.ServiceIdFlag))
	type_ := ctx.String(cmd.GetFlagName(cmd.TypeFlag))
	serviceEndpoint := ctx.String(cmd.GetFlagName(cmd.ServiceEndPointFlag))
	index := ctx.Uint64(cmd.GetFlagName(cmd.IndexFlag))
	txHash, err := ontSdk.Native.OntId.UpdateService(gasPrice, gasLimit, acc, did, []byte(serviceId), []byte(type_), []byte(serviceEndpoint), uint32(index), acc)
	if err != nil {
		return fmt.Errorf("update service err:%s", err)
	}
	fmt.Printf("txHash:%v\n", txHash.ToHexString())
	return nil
}

func checkFileName(ctx *cli.Context) string {
	if ctx.IsSet(cmd.GetFlagName(cmd.WalletFileFlag)) {
		return ctx.String(cmd.GetFlagName(cmd.WalletFileFlag))
	} else {
		//default account file name
		return cmd.DEFAULT_WALLET_FILE_NAME
	}
}

func NewDID(ontSdk *sdk.OntologySdk, acc *sdk.Account, gasPrice, gasLimit uint64) (string, error) {
	did, err := sdk.GenerateID()
	if err != nil {
		return "", err
	}
	err = RegisterDid(did, ontSdk, acc, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	return did, nil
}

func RegisterDid(did string, ontSdk *sdk.OntologySdk, acc *sdk.Account, gasPrice, gasLimit uint64) error {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return fmt.Errorf("ontsdk is nil")
	}
	txHash, err := ontSdk.Native.OntId.RegIDWithPublicKey(gasPrice, gasLimit, acc, did, acc)
	if err != nil {
		return err
	}
	fmt.Printf("Did:  %v,  Hash:%v\n", did, txHash.ToHexString())
	return nil
}

func queryDidDoc(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
	did := ctx.String(cmd.GetFlagName(cmd.DidFlag))
	doc, err := utils.GetDidDocByDid(did, ontSdk)
	if err != nil {
		return fmt.Errorf("err:%s", err)
	}
	fmt.Printf("doc: %v\n", doc)
	return nil
}

func QueryEndPoint(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(ctx.String(cmd.GetFlagName(cmd.RpcUrlFlag)))
	did := ctx.String(cmd.GetFlagName(cmd.DidFlag))
	endPoints, err := utils.GetServiceEndpointByDid(did, ontSdk)
	if err != nil {
		return fmt.Errorf("err:%s", err)
	}
	fmt.Printf("endPoints:%v\n", endPoints)
	return nil
}
