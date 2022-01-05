package application

import (
	"context"
	logicGate "github.com/ariyn/logic-gate"
	"math"
)

func NBitsOr(ctx context.Context, size int) logicGate.Gate {
	truthTables := make(map[int]bool, 0)

	max := int(math.Pow(2, float64(size)))
	for i := 1; i < max; i++ {
		truthTables[i] = true
	}
	truthTables[0] = false

	return logicGate.NewTruthTableGate(ctx, "nBitsOr", size, 1, truthTables)
}
