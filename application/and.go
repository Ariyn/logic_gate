package application

import (
	"context"
	logicGate "github.com/ariyn/logic-gate"
	"math"
)

func NBitsAnd(ctx context.Context, size int) logicGate.Gate {
	truthTables := make(map[int]bool, 0)

	max := int(math.Pow(2, float64(size))) - 1
	for i := 0; i < max; i++ {
		truthTables[i] = false
	}
	truthTables[max] = true

	return logicGate.NewTruthTableGate(ctx, "nBitsAnd", size, 1, truthTables)
}
