package fisco

import (
	"fmt"
	"testing"
)

var fc_pro = &FiscoChainConfig{
	CaFile:     "../conf/fisco_pro/ca.crt",
	Key:        "../conf/fisco_pro/sdk.key",
	Cert:       "../conf/fisco_pro/sdk.crt",
	KeyFile:    "../conf/fisco_pro/kwinin_key_0x5409901f66bcfd12548a164662a1ffa25ca209ea.pem",
	GroupId:    1,
	IpPort:     "120.26.143.5:20200",
	IsHttp:     false,
	ChainId:    1,
	IsSMCrypto: false,
}

var fc_dev = &FiscoChainConfig{
	CaFile:     "../conf/fisco1/ca.crt",
	Key:        "../conf/fisco1/sdk.key",
	Cert:       "../conf/fisco1/sdk.crt",
	KeyFile:    "../conf/fisco1/kwinin_key_0x5409901f66bcfd12548a164662a1ffa25ca209ea.pem",
	GroupId:    1,
	IpPort:     "10.211.55.6:20200",
	IsHttp:     false,
	ChainId:    1,
	IsSMCrypto: false,
}

var fc_dev1 = &FiscoChainConfig{
	CaFile:     "../conf/fisco1/ca.crt",
	Key:        "../conf/fisco1/sdk.key",
	Cert:       "../conf/fisco1/sdk.crt",
	KeyFile:    "../conf/fisco1/test1_key_0xda2abe17b025ba72933b550bfc550cf608da8321.pem",
	GroupId:    1,
	IpPort:     "10.211.55.6:20200",
	IsHttp:     false,
	ChainId:    1,
	IsSMCrypto: false,
}

var fc_dev2 = &FiscoChainConfig{
	CaFile:     "../conf/fisco1/ca.crt",
	Key:        "../conf/fisco1/sdk.key",
	Cert:       "../conf/fisco1/sdk.crt",
	KeyFile:    "../conf/fisco1/test2_key_0x8e1abb21b716cdeb21142b85042419fa7c6ce4e9.pem",
	GroupId:    1,
	IpPort:     "10.211.55.6:20200",
	IsHttp:     false,
	ChainId:    1,
	IsSMCrypto: false,
}

func TestFiscoService_GetBlockNumber(t *testing.T) {
	c := NewFiscoService(fc_dev)
	fmt.Println(c.GetBlockNumber())
}

func TestFiscoService_DeployStore(t *testing.T) {
	c := NewFiscoService(fc_dev1)
	address, tx := c.DeployStore()
	fmt.Printf("address: %s, txhash: %s", address, tx)
}

func TestFiscoService_ProductionGetItems(t *testing.T) {
	c := NewFiscoService(fc_pro)
	address := "0xda6Dd4f54450E08499971F29083B3Ea3718eF1AE"
	pid := "28217b14-33ab-47de-afe5-bea3351e5333"
	r, err := c.ProductionGetItems(address, pid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", r)
}
