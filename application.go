package logic_gate

import "context"

func HalfAdder(ctx context.Context) (hfAdder Gate) {
	xor := XorGate(ctx)
	and := AndGate(ctx)

	inputDigit0 := BasicGate(ctx)
	inputDigit1 := BasicGate(ctx)

	Connect(inputDigit0.Output(0), xor.Input(0))
	Connect(inputDigit1.Output(0), xor.Input(1))
	Connect(inputDigit0.Output(0), and.Input(0))
	Connect(inputDigit1.Output(0), and.Input(1))

	hfAdder = &ComplexGate{
		ctx:        ctx,
		inputSize:  2,
		outputSize: 2,
		inputs:     []Receiver{inputDigit0.Input(0), inputDigit1.Input(0)},
		outputs:    []Transmitter{xor.Output(0), and.Output(0)},
		gates:      []Gate{inputDigit0, inputDigit1, xor, and},
	}

	return
}

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
