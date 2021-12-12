package logic_gate

import "context"

type ComplexGate struct {
	ctx        context.Context
	InputSize  int
	OutputSize int
	inputs     []chan bool
	outputs    []chan bool
	gates      []*Gate
}

func NorGate(ctx context.Context) (g *ComplexGate) {
	orGate := OrGate(ctx)
	notGate := NotGate(ctx)

	notGate.inputs[0] = orGate.outputs[0]

	g = &ComplexGate{
		ctx:        ctx,
		InputSize:  orGate.InputSize,
		OutputSize: notGate.OutputSize,
		inputs:     orGate.inputs,
		outputs:    notGate.outputs,
		gates:      []*Gate{orGate, notGate},
	}

	return
}

func NandGate(ctx context.Context) (g *ComplexGate) {
	andGate := AndGate(ctx)
	notGate := NotGate(ctx)

	notGate.inputs[0] = andGate.outputs[0]

	g = &ComplexGate{
		ctx:        ctx,
		InputSize:  andGate.InputSize,
		OutputSize: notGate.OutputSize,
		inputs:     andGate.inputs,
		outputs:    notGate.outputs,
		gates:      []*Gate{andGate, notGate},
	}

	return
}
