package main

import (
	"github.com/swarleynunez/INVINOS/api"
	"github.com/swarleynunez/INVINOS/core"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// Get program arguments
	deploying, err := strconv.ParseBool(os.Args[1])
	utils.CheckError(err, utils.FatalMode)

	// Initialize and configure node
	core.InitNode(deploying)

	// Execute API
	err = http.ListenAndServe(":"+utils.GetEnv("API_PORT"), api.NewRouter())
	utils.CheckError(err, utils.FatalMode)
}
