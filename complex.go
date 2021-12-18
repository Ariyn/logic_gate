package logic_gate

import "context"

type ComplexGate struct {
	ctx        context.Context
	InputSize  int
	OutputSize int
	inputs     []Transceiver
	outputs    []Transceiver
	gates      []*Gate
}

func NorGate(ctx context.Context) (g *ComplexGate) {
	orGate := OrGate(ctx)
	notGate := NotGate(ctx)

	orGate.outputs[0].outputs = append(orGate.outputs[0].outputs, notGate.inputs[0].input)

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

	andGate.outputs[0].outputs = append(andGate.outputs[0].outputs, notGate.inputs[0].input)

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
