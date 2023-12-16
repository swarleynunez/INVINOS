package types

import (
	"github.com/swarleynunez/INVINOS/core/bindings"
)

type TransitionInfo struct {
	Transition       bindings.TraceabilityTransition
	CurrentProductID uint64
	Quantity         uint64
}
