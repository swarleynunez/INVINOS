package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/swarleynunez/INVINOS/core/utils"
)

var (
	// Unexported and "read-only" global variables
	_ethc  *ethclient.Client
	_ipfsc *rpc.HttpApi

	// Exported and writable global variables
	Instances map[string]string // API token => traceability contract address
)

func InitNode() {

	// Load environment variables
	utils.LoadEnv()

	// Connect to an Ethereum node
	_ethc = ethConnect(utils.GetEnv("ETH_NODE_URL"))

	// Connect to an IPFS node
	_ipfsc = connectToIPFS()

	// Initialize global variables
	Instances = make(map[string]string)

	// Experiments
	//Instances["07a105b1acc4393561498879eac3e4075d2c3e7a48105872ef5d8d5ba9c707fc"] = "0xA5A17A298Cf444Df6e89B391485C3278a8dBdE96"
}
