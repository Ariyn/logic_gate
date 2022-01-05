package application

import (
	"context"
	logicGate "github.com/ariyn/logic-gate"
)

func NorGate(ctx context.Context) (g logicGate.Gate) {
	orGate := logicGate.OrGate(ctx)
	notGate := logicGate.NotGate(ctx)

	logicGate.Connect(orGate.Output(0), notGate.Input(0))

	g = logicGate.NewComplexGate(ctx, "nor gate", orGate.Inputs(), notGate.Outputs(), []logicGate.Gate{orGate, notGate})
	return
}
