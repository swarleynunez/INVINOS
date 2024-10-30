package main

import (
	"github.com/swarleynunez/INVINOS/app"
	"github.com/swarleynunez/INVINOS/core"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
)

func main() {

	// Initialize and configure node
	core.InitNode()

	// Execute APP
	err := http.ListenAndServe(":"+utils.GetEnv("APP_PORT"), app.NewRouter())
	utils.CheckError(err, utils.FatalMode)
}
