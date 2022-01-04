package logic_gate

import "context"

func ComplexFullAdder(ctx context.Context) (g Gate) {
	xor1 := XorGate(ctx)
	xor2 := XorGate(ctx)

	and1 := AndGate(ctx)
	and2 := AndGate(ctx)

	or := OrGate(ctx)

	inputDigit0 := BasicGate(ctx)
	inputDigit1 := BasicGate(ctx)
	inputCarry := BasicGate(ctx)

	Connect(inputDigit0.Output(0), xor1.Input(0))
	Connect(inputDigit1.Output(0), xor1.Input(1))
	Connect(inputDigit0.Output(0), and1.Input(0))
	Connect(inputDigit1.Output(0), and1.Input(1))

	Connect(xor1.Output(0), xor2.Input(0))
	Connect(xor1.Output(0), and2.Input(0))

	Connect(inputCarry.Output(0), xor2.Input(1))
	Connect(inputCarry.Output(0), and2.Input(1))

	Connect(and1.Output(0), or.Input(0))
	Connect(and2.Output(0), or.Input(1))

	g = &ComplexGate{
		ctx:        ctx,
		inputSize:  3,
		outputSize: 2,
		inputs:     []Receiver{inputDigit0.Input(0), inputDigit1.Input(0), inputCarry.Input(0)},
		outputs:    []Transmitter{xor2.Output(0), or.Output(0)},
		gates:      []Gate{inputDigit0, inputDigit1, inputCarry, xor1, xor2, and1, and2, or},
	}

	return
}

func Complex4BitsFullAdder(ctx context.Context) (g Gate) {
	a1 := ComplexFullAdder(ctx)
	a2 := ComplexFullAdder(ctx)
	a3 := ComplexFullAdder(ctx)
	a4 := ComplexFullAdder(ctx)

	Connect(a1.Output(1), a2.Input(2))
	Connect(a2.Output(1), a3.Input(2))
	Connect(a3.Output(1), a4.Input(2))

	g = &ComplexGate{
		ctx:        ctx,
		inputSize:  8,
		outputSize: 5,
		inputs:     []Receiver{a1.Input(0), a2.Input(0), a3.Input(0), a4.Input(0), a1.Input(1), a2.Input(1), a3.Input(1), a4.Input(1)},
		outputs:    []Transmitter{a1.Output(0), a2.Output(0), a3.Output(0), a4.Output(0), a4.Output(1)},
		gates:      []Gate{a1, a2, a3, a4},
	}
	return
}

func Complex4BitsFullSubtractor(ctx context.Context) (g Gate) {
	a1 := ComplexFullAdder(ctx)
	a2 := ComplexFullAdder(ctx)
	a3 := ComplexFullAdder(ctx)
	a4 := ComplexFullAdder(ctx)

	Connect(a1.Output(1), a2.Input(2))
	Connect(a2.Output(1), a3.Input(2))
	Connect(a3.Output(1), a4.Input(2))

	a1.Input(2).Push(true)

	not1 := NotGate(ctx)
	Connect(not1.Output(0), a1.Input(1))

	not2 := NotGate(ctx)
	Connect(not2.Output(0), a2.Input(1))

	not3 := NotGate(ctx)
	Connect(not3.Output(0), a3.Input(1))

	not4 := NotGate(ctx)
	Connect(not4.Output(0), a4.Input(1))

	g = &ComplexGate{
		ctx:        ctx,
		inputSize:  8,
		outputSize: 5,
		inputs:     []Receiver{a1.Input(0), a2.Input(0), a3.Input(0), a4.Input(0), not1.Input(0), not2.Input(0), not3.Input(0), not4.Input(0)},
		outputs:    []Transmitter{a1.Output(0), a2.Output(0), a3.Output(0), a4.Output(0), a4.Output(1)},
		gates:      []Gate{a1, a2, a3, a4, not1, not2, not3, not4},
	}

	return
}
