package fisco

import (
	"backend_ft_tcs/log"
	"context"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

var logger = log.GetLogger("fisco")

type FiscoService interface {
	GetBlockNumber() (int64, error)
	DeployStore() (string, string)
	StoreSetItem(address string)
	StoreItems(address string)

	DeployProduction(pid, version, lot, pname, principal string) (string, string, error)

	ProductionSetItem(address string, pid string, opType int8, opId int64, opName string, value string, principal string) (string, error)

	ProductionGetItems(address string, pid string) (string, error)
}

type fiscoService struct {
	fc     *FiscoChainConfig
	client *client.Client
}

type FiscoChainConfig struct {
	CaFile     string
	Key        string
	Cert       string
	KeyFile    string
	GroupId    int
	IpPort     string
	IsHttp     bool
	ChainId    int64
	IsSMCrypto bool
}

func NewFiscoService(fc *FiscoChainConfig) FiscoService {
	configs, err := conf.ParseConfigOptions(fc.CaFile, fc.Key, fc.Cert, fc.KeyFile, fc.GroupId, fc.IpPort, fc.IsHttp, fc.ChainId, fc.IsSMCrypto)
	if err != nil {
		logger.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial(configs)
	f := new(fiscoService)
	f.fc = fc
	f.client = client
	//return &fiscoService{fc: fc, client: client}
	return f
}

func (f fiscoService) GetBlockNumber() (int64, error) {
	bl, err := f.client.GetBlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return bl, nil
}
