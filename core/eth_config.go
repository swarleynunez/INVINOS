package core

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swarleynunez/INVINOS/core/bindings"
	"github.com/swarleynunez/INVINOS/core/utils"
	"math/big"
	"strconv"
)

func ethConnect(url string) (ethc *ethclient.Client) {

	ethc, err := ethclient.Dial(url)
	utils.CheckError(err, utils.FatalMode)

	_, err = ethc.ChainID(context.Background())
	utils.CheckError(err, utils.FatalMode)

	return
}

func ethTransactor() *bind.TransactOpts {

	// Get account private key
	privKey, err := crypto.HexToECDSA(utils.GetEnv("ETH_ACCOUNT_PRIV_KEY"))
	utils.CheckError(err, utils.FatalMode)

	// Get chain ID (transaction replay protection)
	chainID, err := strconv.ParseUint(utils.GetEnv("ETH_CHAIN_ID"), 10, 64)
	utils.CheckError(err, utils.FatalMode)

	// Transactor using private key
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(chainID)))
	utils.CheckError(err, utils.FatalMode)

	auth.GasLimit = 10485760

	return auth
}

func DeployContracts() string {

	addr, _, _, err := bindings.DeployTraceability(ethTransactor(), _ethc)
	utils.CheckError(err, utils.WarningMode)

	return addr.String()
}

func GetContractInstances(token string) (tinst *bindings.Traceability, einst *bindings.EntityInfo, ainst *bindings.Auth) {

	// Traceability contract instance
	tinst, err := bindings.NewTraceability(common.HexToAddress(Instances[token]), _ethc)
	utils.CheckError(err, utils.WarningMode)

	// EntityInfo contract instance
	eiaddr, err := tinst.EntityInfoContract(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)
	einst, err = bindings.NewEntityInfo(eiaddr, _ethc)
	utils.CheckError(err, utils.WarningMode)

	// Auth contract instance
	aaddr, err := tinst.AuthContract(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)
	ainst, err = bindings.NewAuth(aaddr, _ethc)
	utils.CheckError(err, utils.WarningMode)

	return
}
