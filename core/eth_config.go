package core

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swarleynunez/INVINOS/core/bindings"
	"github.com/swarleynunez/INVINOS/core/utils"
	"math/big"
	"strconv"
)

var (
	ErrMalformedAddr = errors.New("malformed address")
)

func ethConnect(url string) (ethc *ethclient.Client) {

	ethc, err := ethclient.Dial(url)
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

	return auth
}

func DeployContracts() (tinst *bindings.Traceability) {

	taddr, _, tinst, err := bindings.DeployTraceability(ethTransactor(), _ethc)
	utils.CheckError(err, utils.FatalMode)

	// Persist traceability contract address
	utils.SetEnv("TRACEABILITY_CONTRACT", taddr.String())

	return
}

func getTraceabilityInstance() (tinst *bindings.Traceability) {

	// Traceability contract address
	taddr := utils.GetEnv("TRACEABILITY_CONTRACT")

	// Traceability contract instance
	if utils.ValidEthAddress(taddr) {
		i, err := bindings.NewTraceability(common.HexToAddress(taddr), _ethc)
		utils.CheckError(err, utils.FatalMode)
		tinst = i
	} else {
		utils.CheckError(ErrMalformedAddr, utils.FatalMode)
	}

	return
}

func getAuthInstance(tinst *bindings.Traceability) *bindings.Auth {

	// Auth contract address
	aaddr, err := tinst.AuthContract(&bind.CallOpts{})
	utils.CheckError(err, utils.FatalMode)

	// Auth contract instance
	ainst, err := bindings.NewAuth(aaddr, _ethc)
	utils.CheckError(err, utils.FatalMode)

	return ainst
}

func getEntityInfoInstance(tinst *bindings.Traceability) *bindings.EntityInfo {

	// EntityInfo contract address
	eaddr, err := tinst.EntityInfoContract(&bind.CallOpts{})
	utils.CheckError(err, utils.FatalMode)

	// EntityInfo contract instance
	einst, err := bindings.NewEntityInfo(eaddr, _ethc)
	utils.CheckError(err, utils.FatalMode)

	return einst
}
