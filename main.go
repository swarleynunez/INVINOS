package main

import (
	"github.com/swarleynunez/INVINOS/core"
	"os"
)

func main() {

	// Entrypoint
	core.ExecuteServer(os.Args[1])
}
