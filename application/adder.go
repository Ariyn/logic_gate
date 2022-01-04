package application

import (
	"context"
	logicGate "github.com/ariyn/logic-gate"
)

func HalfAdder(ctx context.Context) (hfAdder logicGate.Gate) {
	xor := logicGate.XorGate(ctx)
	and := logicGate.AndGate(ctx)

	inputDigit0 := logicGate.BasicGate(ctx)
	inputDigit1 := logicGate.BasicGate(ctx)

	logicGate.Connect(inputDigit0.Output(0), xor.Input(0))
	logicGate.Connect(inputDigit1.Output(0), xor.Input(1))
	logicGate.Connect(inputDigit0.Output(0), and.Input(0))
	logicGate.Connect(inputDigit1.Output(0), and.Input(1))

	hfAdder = logicGate.NewComplexGate(
		ctx,
		"half adder",
		[]logicGate.Receiver{inputDigit0.Input(0), inputDigit1.Input(0)},
		[]logicGate.Transmitter{xor.Output(0), and.Output(0)},
		[]logicGate.Gate{inputDigit0, inputDigit1, xor, and},
	)

	return
}
