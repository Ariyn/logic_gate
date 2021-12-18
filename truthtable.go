package logic_gate

import (
	"context"
	"sync"
)

var _ Gate = (*TruthTableGate)(nil)

type TruthTableGate struct {
	name           string
	ctx            context.Context
	inputSize      int            // TODO: receiverSize
	inputs         []*Transceiver // TODO: Receiver
	outputSize     int            // TODO: transmitterSize
	outputs        []*Transceiver // TODO: Transmitter
	truthTable     map[int]bool
	previousOutput bool
	handlers       map[HandlerSituation][]gateHandler
	tick           chan *sync.WaitGroup
}

func NewGate(ctx context.Context, inputSize, outputSize int, truthTable map[int]bool) (g Gate) {
	tg := &TruthTableGate{
		ctx:        ctx,
		inputSize:  inputSize,
		outputSize: outputSize,
		inputs:     make([]*Transceiver, inputSize),
		outputs:    make([]*Transceiver, outputSize),
		truthTable: truthTable,
		handlers:   make(map[HandlerSituation][]gateHandler),
		tick:       make(chan *sync.WaitGroup),
	}

	for _, situation := range HandlerSituations {
		tg.handlers[situation] = make([]gateHandler, 0)
	}

	for i := 0; i < tg.inputSize; i++ {
		tg.inputs[i] = NewTransceiver()
	}

	for i := 0; i < tg.outputSize; i++ {
		tg.outputs[i] = NewTransceiver()
	}

	go tg.run()

	return tg
}

func (g *TruthTableGate) InputSize() (size int) {
	return g.inputSize
}

func (g *TruthTableGate) Input(index int) (r Receiver) {
	if g.inputSize < index {
		return nil
	}

	return g.inputs[index]
}

func (g *TruthTableGate) Inputs() (rs []Receiver) {
	for _, r := range g.inputs {
		rs = append(rs, r)
	}

	return
}

func (g *TruthTableGate) OutputSize() int {
	return g.outputSize
}

func (g *TruthTableGate) Output(index int) (t Transmitter) {
	if g.outputSize < index {
		return nil
	}

	return g.outputs[index]
}

func (g *TruthTableGate) Outputs() (ts []Transmitter) {
	for _, r := range g.outputs {
		ts = append(ts, r)
	}

	return
}

func (g *TruthTableGate) Tick(wg *sync.WaitGroup) {
	g.tick <- wg
}

func (g *TruthTableGate) run() {
	defer func() {
		for _, o := range g.outputs {
			o.Close()
		}
	}()

	for {
		if err := g.ctx.Err(); err != nil {
			break
		}

		wg := <-g.tick

		received := make([]bool, g.inputSize)
		state := make([]bool, g.inputSize)

		for i := range g.inputs {
			received[i] = g.inputs[i].receive()
			state[i] = g.inputs[i].status
		}

		if len(g.handlers[AfterInput]) != 0 {
			//for _, f := range g.handlers[AfterInput] {
			//f(g, index, value.Bool())
			//}
		}

		next := g.truthTable[g.getTruthTableIndex(state)]
		if next != g.previousOutput {
			g.outputs[0].transmit(next)
			g.previousOutput = next
		}

		wg.Done()
	}
}

func (g *TruthTableGate) getTruthTableIndex(state []bool) (index int) {
	for i, v := range state {
		index += boolMap[v] << i
	}

	return
}

func (g *TruthTableGate) SetPreviousStatus(status bool) {
	g.previousOutput = status
}
