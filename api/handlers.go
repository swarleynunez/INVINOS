package api

import (
	"github.com/swarleynunez/INVINOS/utils"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Hello"))
	utils.CheckError(err, utils.FatalMode)
}
