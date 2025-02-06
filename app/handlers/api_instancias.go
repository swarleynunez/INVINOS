package handlers

import (
	"github.com/swarleynunez/INVINOS/core"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
)

func CreateInstancePut(w http.ResponseWriter, r *http.Request) {

	// Create a new instance (API token + traceability smart contract)
	token := utils.GenerateAPIToken()
	core.Instances[token] = core.DeployContracts()

	// Successful request
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(token))
	utils.CheckError(err, utils.WarningMode)
}
