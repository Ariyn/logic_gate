package logic_gate

import (
	"context"
	"sync"
)

var _ Gate = (*ComplexGate)(nil)

type ComplexGate struct {
	ctx        context.Context
	inputSize  int
	outputSize int
	inputs     []Receiver
	outputs    []Transmitter
	gates      []Gate
}

func (g *ComplexGate) InputSize() int {
	return g.inputSize
}

func (g *ComplexGate) Inputs() []Receiver {
	return g.inputs
}

func (g *ComplexGate) OutputSize() int {
	return g.outputSize
}

func (g *ComplexGate) Outputs() []Transmitter {
	return g.outputs
}

func (g *ComplexGate) Tick(wg *sync.WaitGroup) {
	// TODO: ComplexGate의 tick은 어떻게 처리?
	// 내부의 게이트들이 각각 engine에 연결?
	// ComplexGate가 직접 뿌려주는 방식?
}

func (g *ComplexGate) Output(index int) (t Transmitter) {
	if g.outputSize < index {
		return nil
	}

	return g.outputs[index]
}

func (g *ComplexGate) Input(index int) (r Receiver) {
	if g.inputSize < index {
		return nil
	}

	return g.inputs[index]
}

func (g *ComplexGate) SetPreviousStatus(status bool) {
	//g.previousOutput = status
}

func NorGate(ctx context.Context) (g Gate) {
	orGate := OrGate(ctx)
	notGate := NotGate(ctx)

	Connect(orGate.Output(0), notGate.Input(0))

	g = &ComplexGate{
		ctx:        ctx,
		inputSize:  orGate.InputSize(),
		outputSize: notGate.OutputSize(),
		inputs:     orGate.Inputs(),
		outputs:    notGate.Outputs(),
		gates:      []Gate{orGate, notGate},
	}

	return
}

func NandGate(ctx context.Context) (g Gate) {
	andGate := AndGate(ctx)
	notGate := NotGate(ctx)

	Connect(andGate.Output(0), notGate.Input(0))

	g = &ComplexGate{
		ctx:        ctx,
		inputSize:  andGate.InputSize(),
		outputSize: notGate.OutputSize(),
		inputs:     andGate.Inputs(),
		outputs:    notGate.Outputs(),
		gates:      []Gate{andGate, notGate},
	}

	return
}

func FlipFlopSR(ctx context.Context) (g Gate) {
	rGate := NorGate(ctx)
	sGate := NorGate(ctx)

	Connect(rGate.Output(0), sGate.Input(1))
	Connect(sGate.Output(0), rGate.Input(1))

	rGate.Input(1).SetCurrentStatus(true)
	rGate.(*ComplexGate).gates[0].SetPreviousStatus(false)
	rGate.(*ComplexGate).gates[1].SetPreviousStatus(true)
	sGate.(*ComplexGate).gates[0].SetPreviousStatus(true)

	return &ComplexGate{
		ctx:        ctx,
		inputSize:  2,
		outputSize: 2,
		inputs: []Receiver{
			rGate.Input(0),
			sGate.Input(0),
		},
		outputs: []Transmitter{
			rGate.Output(0),
			sGate.Output(0),
		},
		gates: []Gate{
			rGate, sGate,
		},
	}
}
