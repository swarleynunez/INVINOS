package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/swarleynunez/INVINOS/core/bindings"
	"github.com/swarleynunez/INVINOS/core/utils"
)

var (
	// Unexported and "read-only" global variables
	_ethc  *ethclient.Client
	_tinst *bindings.Traceability
	_ainst *bindings.Auth
	_einst *bindings.EntityInfo
	_ipfsc *rpc.HttpApi
)

func InitNode(deploying bool) {

	// Load environment variables
	utils.LoadEnv()

	// Connect to an Ethereum node
	_ethc = ethConnect(utils.GetEnv("ETH_NODE_URL"))

	// Deploy contracts and get instances
	if deploying {
		_tinst = DeployContracts()
	} else {
		_tinst = getTraceabilityInstance()
	}
	_ainst = getAuthInstance(_tinst)
	_einst = getEntityInfoInstance(_tinst)

	// Connect to an IPFS node
	_ipfsc = connectToIPFS()
}

func ReloadContractInstances() {

	_tinst = getTraceabilityInstance()
	_ainst = getAuthInstance(_tinst)
	_einst = getEntityInfoInstance(_tinst)
}
