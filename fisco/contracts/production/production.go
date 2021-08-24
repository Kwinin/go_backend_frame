// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package production

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// productionItem is an auto generated low-level Go binding around an user-defined struct.
type productionItem struct {
	OpName    string
	OpType    int8
	OpId      int64
	ItemValue string
	IsVaild   bool
	Principal string
}

// ProductionABI is the input ABI used to generate the binding from.
const ProductionABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_v\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_lot\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_principal\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_opType\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"opId\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"_opType\",\"type\":\"int8\"},{\"internalType\":\"int64\",\"name\":\"_opId\",\"type\":\"int64\"}],\"name\":\"getItem\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"opName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"itemValue\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"principal\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"}],\"name\":\"getItems\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"opName\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"opType\",\"type\":\"int8\"},{\"internalType\":\"int64\",\"name\":\"opId\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"itemValue\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"principal\",\"type\":\"string\"}],\"internalType\":\"structproduction.Item[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"}],\"name\":\"getOpIds\",\"outputs\":[{\"internalType\":\"int64[]\",\"name\":\"opIds\",\"type\":\"int64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"}],\"name\":\"getOpTypes\",\"outputs\":[{\"internalType\":\"int8[]\",\"name\":\"opTypes\",\"type\":\"int8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"},{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"name\":\"itemMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"opName\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"opType\",\"type\":\"int8\"},{\"internalType\":\"int64\",\"name\":\"opId\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"itemValue\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVaild\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"principal\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pname\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"principal\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pid\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"_opType\",\"type\":\"int8\"},{\"internalType\":\"int64\",\"name\":\"_opId\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"_opName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"principal\",\"type\":\"string\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProductionBin is the compiled bytecode used for deploying new contracts.
var ProductionBin = "0x60806040523480156200001157600080fd5b506040516200218e3803806200218e8339818101604052810190620000379190620001b1565b83600090805190602001906200004f929190620000a6565b50846001908051906020019062000068929190620000a6565b50816002908051906020019062000081929190620000a6565b5080600390805190602001906200009a929190620000a6565b50505050505062000341565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000e957805160ff19168380011785556200011a565b828001600101855582156200011a579182015b8281111562000119578251825591602001919060010190620000fc565b5b5090506200012991906200012d565b5090565b6200015291905b808211156200014e57600081600090555060010162000134565b5090565b90565b600082601f8301126200016757600080fd5b81516200017e6200017882620002de565b620002b0565b915080825260208301602083018583830111156200019b57600080fd5b620001a88382846200030b565b50505092915050565b600080600080600060a08688031215620001ca57600080fd5b600086015167ffffffffffffffff811115620001e557600080fd5b620001f38882890162000155565b955050602086015167ffffffffffffffff8111156200021157600080fd5b6200021f8882890162000155565b945050604086015167ffffffffffffffff8111156200023d57600080fd5b6200024b8882890162000155565b935050606086015167ffffffffffffffff8111156200026957600080fd5b620002778882890162000155565b925050608086015167ffffffffffffffff8111156200029557600080fd5b620002a38882890162000155565b9150509295509295909350565b6000604051905081810181811067ffffffffffffffff82111715620002d457600080fd5b8060405250919050565b600067ffffffffffffffff821115620002f657600080fd5b601f19601f8301169050602081019050919050565b60005b838110156200032b5780820151818401526020810190506200030e565b838111156200033b576000848401525b50505050565b611e3d80620003516000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639a77b665116100665780639a77b66514610172578063ba5d3078146101a7578063cb78b527146101c5578063d6894369146101f5578063f1068454146102115761009e565b806322664e65146100a35780633cb25546146100c157806354c424e1146100f157806354fd4d50146101215780637861d1c51461013f575b600080fd5b6100ab61022f565b6040516100b89190611aad565b60405180910390f35b6100db60048036038101906100d69190611547565b6102cd565b6040516100e89190611a69565b60405180910390f35b61010b60048036038101906101069190611547565b610371565b6040516101189190611a47565b60405180910390f35b610129610415565b6040516101369190611aad565b60405180910390f35b61015960048036038101906101549190611588565b6104b3565b6040516101699493929190611b98565b60405180910390f35b61018c60048036038101906101879190611588565b61077c565b60405161019e96959493929190611b22565b60405180910390f35b6101af6109d7565b6040516101bc9190611aad565b60405180910390f35b6101df60048036038101906101da9190611547565b610a75565b6040516101ec9190611a8b565b60405180910390f35b61020f600480360381019061020a91906115ef565b610f10565b005b610219611349565b6040516102269190611aad565b60405180910390f35b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c55780601f1061029a576101008083540402835291602001916102c5565b820191906000526020600020905b8154815290600101906020018083116102a857829003601f168201915b505050505081565b6060806005836040516102e09190611a30565b908152602001604051809103902060000180548060200260200160405190810160405280929190818152602001828054801561036157602002820191906000526020600020906000905b82829054906101000a900460000b60000b8152602001906001019060208260000104928301926001038202915080841161032a5790505b5050505050905080915050919050565b6060806005836040516103849190611a30565b908152602001604051809103902060010180548060200260200160405190810160405280929190818152602001828054801561040557602002820191906000526020600020906000905b82829054906101000a900460070b60070b815260200190600801906020826007010492830192600103820291508084116103ce5790505b5050505050905080915050919050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104ab5780601f10610480576101008083540402835291602001916104ab565b820191906000526020600020905b81548152906001019060200180831161048e57829003601f168201915b505050505081565b606080600060606104c26113e7565b6004886040516104d29190611a30565b908152602001604051809103902060008860000b60000b815260200190815260200160002060008760070b60070b81526020019081526020016000206040518060c0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105b25780601f10610587576101008083540402835291602001916105b2565b820191906000526020600020905b81548152906001019060200180831161059557829003601f168201915b505050505081526020016001820160009054906101000a900460000b60000b60000b81526020016001820160019054906101000a900460070b60070b60070b8152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561068e5780601f106106635761010080835404028352916020019161068e565b820191906000526020600020905b81548152906001019060200180831161067157829003601f168201915b505050505081526020016003820160009054906101000a900460ff16151515158152602001600482018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561074b5780601f106107205761010080835404028352916020019161074b565b820191906000526020600020905b81548152906001019060200180831161072e57829003601f168201915b50505050508152505090508060000151816060015182608001518360a0015194509450945094505093509350935093565b6004838051602081018201805184825260208301602085012081835280955050505050506020528160005260406000206020528060005260406000206000925092505050806000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108585780601f1061082d57610100808354040283529160200191610858565b820191906000526020600020905b81548152906001019060200180831161083b57829003601f168201915b5050505050908060010160009054906101000a900460000b908060010160019054906101000a900460070b90806002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561091c5780601f106108f15761010080835404028352916020019161091c565b820191906000526020600020905b8154815290600101906020018083116108ff57829003601f168201915b5050505050908060030160009054906101000a900460ff1690806004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109cd5780601f106109a2576101008083540402835291602001916109cd565b820191906000526020600020905b8154815290600101906020018083116109b057829003601f168201915b5050505050905086565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6d5780601f10610a4257610100808354040283529160200191610a6d565b820191906000526020600020905b815481529060010190602001808311610a5057829003601f168201915b505050505081565b606080600583604051610a889190611a30565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020018280548015610b0957602002820191906000526020600020906000905b82829054906101000a900460070b60070b81526020019060080190602082600701049283019260010382029150808411610ad25790505b505050505090506060600584604051610b229190611a30565b9081526020016040518091039020600001805480602002602001604051908101604052809291908181526020018280548015610ba357602002820191906000526020600020906000905b82829054906101000a900460000b60000b81526020019060010190602082600001049283019260010382029150808411610b6c5790505b505050505090506060825167ffffffffffffffff81118015610bc457600080fd5b50604051908082528060200260200182016040528015610bfe57816020015b610beb6113e7565b815260200190600190039081610be35790505b50905060008090505b8351811015610f04576000848281518110610c1e57fe5b602002602001015190506000848381518110610c3657fe5b60200260200101519050610c486113e7565b600489604051610c589190611a30565b908152602001604051809103902060008360000b60000b815260200190815260200160002060008460070b60070b81526020019081526020016000206040518060c0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d385780601f10610d0d57610100808354040283529160200191610d38565b820191906000526020600020905b815481529060010190602001808311610d1b57829003601f168201915b505050505081526020016001820160009054906101000a900460000b60000b60000b81526020016001820160019054906101000a900460070b60070b60070b8152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e145780601f10610de957610100808354040283529160200191610e14565b820191906000526020600020905b815481529060010190602001808311610df757829003601f168201915b505050505081526020016003820160009054906101000a900460ff16151515158152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ed15780601f10610ea657610100808354040283529160200191610ed1565b820191906000526020600020905b815481529060010190602001808311610eb457829003601f168201915b505050505081525050905080858581518110610ee957fe5b60200260200101819052505050508080600101915050610c07565b50809350505050919050565b600486604051610f209190611a30565b908152602001604051809103902060008660000b60000b815260200190815260200160002060008560070b60070b815260200190815260200160002060030160009054906101000a900460ff1615610fad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa490611bf2565b60405180910390fd5b600586604051610fbd9190611a30565b90815260200160405180910390206000018590806001815401808255809150506001900390600052602060002090602091828204019190069091909190916101000a81548160ff021916908360000b60ff1602179055506005866040516110249190611a30565b90815260200160405180910390206001018490806001815401808255809150506001900390600052602060002090600491828204019190066008029091909190916101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff160217905550600160048760405161109e9190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b815260200190815260200160002060030160006101000a81548160ff021916908315150217905550826004876040516111059190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b81526020019081526020016000206000019080519060200190611157929190611424565b50846004876040516111699190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b815260200190815260200160002060010160006101000a81548160ff021916908360000b60ff160217905550836004876040516111d49190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b815260200190815260200160002060010160016101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff1602179055508160048760405161124d9190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b8152602001908152602001600020600201908051906020019061129f929190611424565b50806004876040516112b19190611a30565b908152602001604051809103902060008760000b60000b815260200190815260200160002060008660070b60070b81526020019081526020016000206004019080519060200190611303929190611424565b507f4535fe6856bdfb1b2080fbfd90aa254a40ff0ea8786cd7044818533eaed2a94b868686856040516113399493929190611acf565b60405180910390a1505050505050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113df5780601f106113b4576101008083540402835291602001916113df565b820191906000526020600020905b8154815290600101906020018083116113c257829003601f168201915b505050505081565b6040518060c00160405280606081526020016000800b8152602001600060070b815260200160608152602001600015158152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061146557805160ff1916838001178555611493565b82800160010185558215611493579182015b82811115611492578251825591602001919060010190611477565b5b5090506114a091906114a4565b5090565b6114c691905b808211156114c25760008160009055506001016114aa565b5090565b90565b6000813590506114d881611dd9565b92915050565b6000813590506114ed81611df0565b92915050565b600082601f83011261150457600080fd5b813561151761151282611c3f565b611c12565b9150808252602083016020830185838301111561153357600080fd5b61153e838284611d86565b50505092915050565b60006020828403121561155957600080fd5b600082013567ffffffffffffffff81111561157357600080fd5b61157f848285016114f3565b91505092915050565b60008060006060848603121561159d57600080fd5b600084013567ffffffffffffffff8111156115b757600080fd5b6115c3868287016114f3565b93505060206115d4868287016114de565b92505060406115e5868287016114c9565b9150509250925092565b60008060008060008060c0878903121561160857600080fd5b600087013567ffffffffffffffff81111561162257600080fd5b61162e89828a016114f3565b965050602061163f89828a016114de565b955050604061165089828a016114c9565b945050606087013567ffffffffffffffff81111561166d57600080fd5b61167989828a016114f3565b935050608087013567ffffffffffffffff81111561169657600080fd5b6116a289828a016114f3565b92505060a087013567ffffffffffffffff8111156116bf57600080fd5b6116cb89828a016114f3565b9150509295509295509295565b60006116e4838361186b565b60208301905092915050565b60006116fc8383611898565b60208301905092915050565b60006117148383611999565b905092915050565b600061172782611c9b565b6117318185611cee565b935061173c83611c6b565b8060005b8381101561176d57815161175488826116d8565b975061175f83611cc7565b925050600181019050611740565b5085935050505092915050565b600061178582611ca6565b61178f8185611cff565b935061179a83611c7b565b8060005b838110156117cb5781516117b288826116f0565b97506117bd83611cd4565b92505060018101905061179e565b5085935050505092915050565b60006117e382611cb1565b6117ed8185611d10565b9350836020820285016117ff85611c8b565b8060005b8581101561183b578484038952815161181c8582611708565b945061182783611ce1565b925060208a01995050600181019050611803565b50829750879550505050505092915050565b61185681611d4e565b82525050565b61186581611d4e565b82525050565b61187481611d5a565b82525050565b61188381611d5a565b82525050565b61189281611d74565b82525050565b6118a181611d67565b82525050565b6118b081611d67565b82525050565b60006118c182611cbc565b6118cb8185611d21565b93506118db818560208601611d95565b6118e481611dc8565b840191505092915050565b60006118fa82611cbc565b6119048185611d32565b9350611914818560208601611d95565b61191d81611dc8565b840191505092915050565b600061193382611cbc565b61193d8185611d43565b935061194d818560208601611d95565b80840191505092915050565b6000611966601883611d32565b91507f63616e6e277420736574206974656d4d617020616761696e00000000000000006000830152602082019050919050565b600060c08301600083015184820360008601526119b682826118b6565b91505060208301516119cb6020860182611898565b5060408301516119de604086018261186b565b50606083015184820360608601526119f682826118b6565b9150506080830151611a0b608086018261184d565b5060a083015184820360a0860152611a2382826118b6565b9150508091505092915050565b6000611a3c8284611928565b915081905092915050565b60006020820190508181036000830152611a61818461171c565b905092915050565b60006020820190508181036000830152611a83818461177a565b905092915050565b60006020820190508181036000830152611aa581846117d8565b905092915050565b60006020820190508181036000830152611ac781846118ef565b905092915050565b60006080820190508181036000830152611ae981876118ef565b9050611af86020830186611889565b611b05604083018561187a565b8181036060830152611b1781846118ef565b905095945050505050565b600060c0820190508181036000830152611b3c81896118ef565b9050611b4b60208301886118a7565b611b58604083018761187a565b8181036060830152611b6a81866118ef565b9050611b79608083018561185c565b81810360a0830152611b8b81846118ef565b9050979650505050505050565b60006080820190508181036000830152611bb281876118ef565b90508181036020830152611bc681866118ef565b9050611bd5604083018561185c565b8181036060830152611be781846118ef565b905095945050505050565b60006020820190508181036000830152611c0b81611959565b9050919050565b6000604051905081810181811067ffffffffffffffff82111715611c3557600080fd5b8060405250919050565b600067ffffffffffffffff821115611c5657600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008115159050919050565b60008160070b9050919050565b60008160000b9050919050565b6000611d7f82611d67565b9050919050565b82818337600083830152505050565b60005b83811015611db3578082015181840152602081019050611d98565b83811115611dc2576000848401525b50505050565b6000601f19601f8301169050919050565b611de281611d5a565b8114611ded57600080fd5b50565b611df981611d67565b8114611e0457600080fd5b5056fea26469706673582212202b21b3f7844c53a221022fddd200611bd59c20df8975784501ae328b3a1fe5a764736f6c634300060a0033"

// DeployProduction deploys a new Ethereum contract, binding an instance of Production to it.
func DeployProduction(auth *bind.TransactOpts, backend bind.ContractBackend, _pid string, _v string, _lot string, _pname string, _principal string) (common.Address, *types.Transaction, *Production, error) {
	parsed, err := abi.JSON(strings.NewReader(ProductionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProductionBin), backend, _pid, _v, _lot, _pname, _principal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Production{ProductionCaller: ProductionCaller{contract: contract}, ProductionTransactor: ProductionTransactor{contract: contract}, ProductionFilterer: ProductionFilterer{contract: contract}}, nil
}

// Production is an auto generated Go binding around an Ethereum contract.
type Production struct {
	ProductionCaller     // Read-only binding to the contract
	ProductionTransactor // Write-only binding to the contract
	ProductionFilterer   // Log filterer for contract events
}

// ProductionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProductionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProductionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProductionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProductionSession struct {
	Contract     *Production       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProductionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProductionCallerSession struct {
	Contract *ProductionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProductionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProductionTransactorSession struct {
	Contract     *ProductionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProductionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProductionRaw struct {
	Contract *Production // Generic contract binding to access the raw methods on
}

// ProductionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProductionCallerRaw struct {
	Contract *ProductionCaller // Generic read-only contract binding to access the raw methods on
}

// ProductionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProductionTransactorRaw struct {
	Contract *ProductionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProduction creates a new instance of Production, bound to a specific deployed contract.
func NewProduction(address common.Address, backend bind.ContractBackend) (*Production, error) {
	contract, err := bindProduction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Production{ProductionCaller: ProductionCaller{contract: contract}, ProductionTransactor: ProductionTransactor{contract: contract}, ProductionFilterer: ProductionFilterer{contract: contract}}, nil
}

// NewProductionCaller creates a new read-only instance of Production, bound to a specific deployed contract.
func NewProductionCaller(address common.Address, caller bind.ContractCaller) (*ProductionCaller, error) {
	contract, err := bindProduction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProductionCaller{contract: contract}, nil
}

// NewProductionTransactor creates a new write-only instance of Production, bound to a specific deployed contract.
func NewProductionTransactor(address common.Address, transactor bind.ContractTransactor) (*ProductionTransactor, error) {
	contract, err := bindProduction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProductionTransactor{contract: contract}, nil
}

// NewProductionFilterer creates a new log filterer instance of Production, bound to a specific deployed contract.
func NewProductionFilterer(address common.Address, filterer bind.ContractFilterer) (*ProductionFilterer, error) {
	contract, err := bindProduction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProductionFilterer{contract: contract}, nil
}

// bindProduction binds a generic wrapper to an already deployed contract.
func bindProduction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProductionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Production *ProductionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Production.Contract.ProductionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Production *ProductionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.ProductionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Production *ProductionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.ProductionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Production *ProductionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Production.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Production *ProductionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Production *ProductionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.contract.Transact(opts, method, params...)
}

// GetItem is a free data retrieval call binding the contract method 0x7861d1c5.
//
// Solidity: function getItem(string _pid, int8 _opType, int64 _opId) constant returns(string opName, string itemValue, bool isVaild, string principal)
func (_Production *ProductionCaller) GetItem(opts *bind.CallOpts, _pid string, _opType int8, _opId int64) (struct {
	OpName    string
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	ret := new(struct {
		OpName    string
		ItemValue string
		IsVaild   bool
		Principal string
	})
	out := ret
	err := _Production.contract.Call(opts, out, "getItem", _pid, _opType, _opId)
	return *ret, err
}

// GetItem is a free data retrieval call binding the contract method 0x7861d1c5.
//
// Solidity: function getItem(string _pid, int8 _opType, int64 _opId) constant returns(string opName, string itemValue, bool isVaild, string principal)
func (_Production *ProductionSession) GetItem(_pid string, _opType int8, _opId int64) (struct {
	OpName    string
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	return _Production.Contract.GetItem(&_Production.CallOpts, _pid, _opType, _opId)
}

// GetItem is a free data retrieval call binding the contract method 0x7861d1c5.
//
// Solidity: function getItem(string _pid, int8 _opType, int64 _opId) constant returns(string opName, string itemValue, bool isVaild, string principal)
func (_Production *ProductionCallerSession) GetItem(_pid string, _opType int8, _opId int64) (struct {
	OpName    string
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	return _Production.Contract.GetItem(&_Production.CallOpts, _pid, _opType, _opId)
}

// GetItems is a free data retrieval call binding the contract method 0xcb78b527.
//
// Solidity: function getItems(string _pid) constant returns([]productionItem items)
func (_Production *ProductionCaller) GetItems(opts *bind.CallOpts, _pid string) ([]productionItem, error) {
	var (
		ret0 = new([]productionItem)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "getItems", _pid)
	return *ret0, err
}

// GetItems is a free data retrieval call binding the contract method 0xcb78b527.
//
// Solidity: function getItems(string _pid) constant returns([]productionItem items)
func (_Production *ProductionSession) GetItems(_pid string) ([]productionItem, error) {
	return _Production.Contract.GetItems(&_Production.CallOpts, _pid)
}

// GetItems is a free data retrieval call binding the contract method 0xcb78b527.
//
// Solidity: function getItems(string _pid) constant returns([]productionItem items)
func (_Production *ProductionCallerSession) GetItems(_pid string) ([]productionItem, error) {
	return _Production.Contract.GetItems(&_Production.CallOpts, _pid)
}

// GetOpIds is a free data retrieval call binding the contract method 0x54c424e1.
//
// Solidity: function getOpIds(string _pid) constant returns(int64[] opIds)
func (_Production *ProductionCaller) GetOpIds(opts *bind.CallOpts, _pid string) ([]int64, error) {
	var (
		ret0 = new([]int64)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "getOpIds", _pid)
	return *ret0, err
}

// GetOpIds is a free data retrieval call binding the contract method 0x54c424e1.
//
// Solidity: function getOpIds(string _pid) constant returns(int64[] opIds)
func (_Production *ProductionSession) GetOpIds(_pid string) ([]int64, error) {
	return _Production.Contract.GetOpIds(&_Production.CallOpts, _pid)
}

// GetOpIds is a free data retrieval call binding the contract method 0x54c424e1.
//
// Solidity: function getOpIds(string _pid) constant returns(int64[] opIds)
func (_Production *ProductionCallerSession) GetOpIds(_pid string) ([]int64, error) {
	return _Production.Contract.GetOpIds(&_Production.CallOpts, _pid)
}

// GetOpTypes is a free data retrieval call binding the contract method 0x3cb25546.
//
// Solidity: function getOpTypes(string _pid) constant returns(int8[] opTypes)
func (_Production *ProductionCaller) GetOpTypes(opts *bind.CallOpts, _pid string) ([]int8, error) {
	var (
		ret0 = new([]int8)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "getOpTypes", _pid)
	return *ret0, err
}

// GetOpTypes is a free data retrieval call binding the contract method 0x3cb25546.
//
// Solidity: function getOpTypes(string _pid) constant returns(int8[] opTypes)
func (_Production *ProductionSession) GetOpTypes(_pid string) ([]int8, error) {
	return _Production.Contract.GetOpTypes(&_Production.CallOpts, _pid)
}

// GetOpTypes is a free data retrieval call binding the contract method 0x3cb25546.
//
// Solidity: function getOpTypes(string _pid) constant returns(int8[] opTypes)
func (_Production *ProductionCallerSession) GetOpTypes(_pid string) ([]int8, error) {
	return _Production.Contract.GetOpTypes(&_Production.CallOpts, _pid)
}

// ItemMap is a free data retrieval call binding the contract method 0x9a77b665.
//
// Solidity: function itemMap(string , int8 , int64 ) constant returns(string opName, int8 opType, int64 opId, string itemValue, bool isVaild, string principal)
func (_Production *ProductionCaller) ItemMap(opts *bind.CallOpts, arg0 string, arg1 int8, arg2 int64) (struct {
	OpName    string
	OpType    int8
	OpId      int64
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	ret := new(struct {
		OpName    string
		OpType    int8
		OpId      int64
		ItemValue string
		IsVaild   bool
		Principal string
	})
	out := ret
	err := _Production.contract.Call(opts, out, "itemMap", arg0, arg1, arg2)
	return *ret, err
}

// ItemMap is a free data retrieval call binding the contract method 0x9a77b665.
//
// Solidity: function itemMap(string , int8 , int64 ) constant returns(string opName, int8 opType, int64 opId, string itemValue, bool isVaild, string principal)
func (_Production *ProductionSession) ItemMap(arg0 string, arg1 int8, arg2 int64) (struct {
	OpName    string
	OpType    int8
	OpId      int64
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	return _Production.Contract.ItemMap(&_Production.CallOpts, arg0, arg1, arg2)
}

// ItemMap is a free data retrieval call binding the contract method 0x9a77b665.
//
// Solidity: function itemMap(string , int8 , int64 ) constant returns(string opName, int8 opType, int64 opId, string itemValue, bool isVaild, string principal)
func (_Production *ProductionCallerSession) ItemMap(arg0 string, arg1 int8, arg2 int64) (struct {
	OpName    string
	OpType    int8
	OpId      int64
	ItemValue string
	IsVaild   bool
	Principal string
}, error) {
	return _Production.Contract.ItemMap(&_Production.CallOpts, arg0, arg1, arg2)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() constant returns(string)
func (_Production *ProductionCaller) Pid(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "pid")
	return *ret0, err
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() constant returns(string)
func (_Production *ProductionSession) Pid() (string, error) {
	return _Production.Contract.Pid(&_Production.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() constant returns(string)
func (_Production *ProductionCallerSession) Pid() (string, error) {
	return _Production.Contract.Pid(&_Production.CallOpts)
}

// Pname is a free data retrieval call binding the contract method 0x22664e65.
//
// Solidity: function pname() constant returns(string)
func (_Production *ProductionCaller) Pname(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "pname")
	return *ret0, err
}

// Pname is a free data retrieval call binding the contract method 0x22664e65.
//
// Solidity: function pname() constant returns(string)
func (_Production *ProductionSession) Pname() (string, error) {
	return _Production.Contract.Pname(&_Production.CallOpts)
}

// Pname is a free data retrieval call binding the contract method 0x22664e65.
//
// Solidity: function pname() constant returns(string)
func (_Production *ProductionCallerSession) Pname() (string, error) {
	return _Production.Contract.Pname(&_Production.CallOpts)
}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() constant returns(string)
func (_Production *ProductionCaller) Principal(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "principal")
	return *ret0, err
}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() constant returns(string)
func (_Production *ProductionSession) Principal() (string, error) {
	return _Production.Contract.Principal(&_Production.CallOpts)
}

// Principal is a free data retrieval call binding the contract method 0xba5d3078.
//
// Solidity: function principal() constant returns(string)
func (_Production *ProductionCallerSession) Principal() (string, error) {
	return _Production.Contract.Principal(&_Production.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Production *ProductionCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Production.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Production *ProductionSession) Version() (string, error) {
	return _Production.Contract.Version(&_Production.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Production *ProductionCallerSession) Version() (string, error) {
	return _Production.Contract.Version(&_Production.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xd6894369.
//
// Solidity: function setItem(string _pid, int8 _opType, int64 _opId, string _opName, string _value, string principal) returns()
func (_Production *ProductionTransactor) SetItem(opts *bind.TransactOpts, _pid string, _opType int8, _opId int64, _opName string, _value string, principal string) (*types.Transaction, *types.Receipt, error) {
	return _Production.contract.Transact(opts, "setItem", _pid, _opType, _opId, _opName, _value, principal)
}

// SetItem is a paid mutator transaction binding the contract method 0xd6894369.
//
// Solidity: function setItem(string _pid, int8 _opType, int64 _opId, string _opName, string _value, string principal) returns()
func (_Production *ProductionSession) SetItem(_pid string, _opType int8, _opId int64, _opName string, _value string, principal string) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.SetItem(&_Production.TransactOpts, _pid, _opType, _opId, _opName, _value, principal)
}

// SetItem is a paid mutator transaction binding the contract method 0xd6894369.
//
// Solidity: function setItem(string _pid, int8 _opType, int64 _opId, string _opName, string _value, string principal) returns()
func (_Production *ProductionTransactorSession) SetItem(_pid string, _opType int8, _opId int64, _opName string, _value string, principal string) (*types.Transaction, *types.Receipt, error) {
	return _Production.Contract.SetItem(&_Production.TransactOpts, _pid, _opType, _opId, _opName, _value, principal)
}

// ProductionItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Production contract.
type ProductionItemSetIterator struct {
	Event *ProductionItemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProductionItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductionItemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProductionItemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProductionItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductionItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductionItemSet represents a ItemSet event raised by the Production contract.
type ProductionItemSet struct {
	Pid    string
	OpType *big.Int
	OpId   int64
	Value  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x4535fe6856bdfb1b2080fbfd90aa254a40ff0ea8786cd7044818533eaed2a94b.
//
// Solidity: event ItemSet(string pid, int256 _opType, int64 opId, string value)
func (_Production *ProductionFilterer) FilterItemSet(opts *bind.FilterOpts) (*ProductionItemSetIterator, error) {

	logs, sub, err := _Production.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &ProductionItemSetIterator{contract: _Production.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x4535fe6856bdfb1b2080fbfd90aa254a40ff0ea8786cd7044818533eaed2a94b.
//
// Solidity: event ItemSet(string pid, int256 _opType, int64 opId, string value)
func (_Production *ProductionFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *ProductionItemSet) (event.Subscription, error) {

	logs, sub, err := _Production.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductionItemSet)
				if err := _Production.contract.UnpackLog(event, "ItemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseItemSet is a log parse operation binding the contract event 0x4535fe6856bdfb1b2080fbfd90aa254a40ff0ea8786cd7044818533eaed2a94b.
//
// Solidity: event ItemSet(string pid, int256 _opType, int64 opId, string value)
func (_Production *ProductionFilterer) ParseItemSet(log types.Log) (*ProductionItemSet, error) {
	event := new(ProductionItemSet)
	if err := _Production.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
