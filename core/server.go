package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/swarleynunez/INVINOS/api"
	"github.com/swarleynunez/INVINOS/core/utils"
	"net/http"
)

func ExecuteServer(port string) {

	// HTTP router
	r := chi.NewRouter()

	// Set API routes
	api.Routes(r)

	// Execute server
	err := http.ListenAndServe(":"+port, r)
	utils.CheckError(err, utils.FatalMode)
}
