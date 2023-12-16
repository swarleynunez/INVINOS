package types

import (
	"github.com/swarleynunez/INVINOS/core/bindings"
	"math/big"
)

type Product struct {
	Quantity    *big.Int
	Tv          bindings.TraceabilityTraceabilityVector
	Partitioned bool
	Completed   bool
}
