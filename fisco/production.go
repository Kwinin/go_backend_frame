package fisco

import (
	"backend_ft_tcs/fisco/contracts/production"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

func (f fiscoService) DeployProduction(pid, version, lot, pname, principal string) (string, string, error) {
	address, tx, _, err := production.DeployProduction(f.client.GetTransactOpts(), f.client, pid, version, lot, pname, principal)
	if err != nil {
		return "", "", err
	}

	return address.Hex(), tx.Hash().Hex(), nil
}

func (f fiscoService) ProductionSetItem(address string, pid string, opType int8, opId int64, opName string, value string, principal string) (string, error) {
	contractAddress := common.HexToAddress(address)
	instance, err := production.NewProduction(contractAddress, f.client)
	if err != nil {
		return "", err
	}
	productionSession := &production.ProductionSession{Contract: instance, CallOpts: *f.client.GetCallOpts(), TransactOpts: *f.client.GetTransactOpts()}
	tx, _, err := productionSession.SetItem(pid, opType, opId, opName, value, principal)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (f fiscoService) ProductionGetItems(address string, pid string) (string, error) {
	contractAddress := common.HexToAddress(address)
	instance, err := production.NewProduction(contractAddress, f.client)
	if err != nil {
		return "", err
	}
	productionSession := &production.ProductionSession{Contract: instance, CallOpts: *f.client.GetCallOpts(), TransactOpts: *f.client.GetTransactOpts()}
	proItems, err := productionSession.GetItems(pid)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(proItems)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
