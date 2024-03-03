package handlers

import (
	"github.com/swarleynunez/INVINOS/core"
	"net/http"
)

func DeployContractsPut(w http.ResponseWriter, r *http.Request) {

	// Deploy new smart contracts
	_ = core.DeployContracts()

	// Reload smart contract instances
	core.ReloadContractInstances()

	// Successful request
	w.WriteHeader(http.StatusOK)
}
