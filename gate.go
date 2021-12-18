package logic_gate

import (
	"context"
	"sync"
)

var boolMap = map[bool]int{
	false: 0,
	true:  1,
}

type HandlerSituation int

const (
	AfterInput HandlerSituation = iota + 1
)

var HandlerSituations = []HandlerSituation{
	AfterInput,
}

type gateHandler func(g *Gate, index int, input bool)

// TODO: rename Gate to `TruthTableGate`
// opposite is `ComplexGate`
// Gate should be interface
type Gate struct {
	name           string
	ctx            context.Context
	wg             *sync.WaitGroup
	InputSize      int
	OutputSize     int
	inputs         []Transceiver
	outputs        []Transceiver
	truthTable     map[int]bool
	previousOutput bool
	handlers       map[HandlerSituation][]gateHandler
	tick           chan *sync.WaitGroup
}

func NewGate(ctx context.Context, inputSize, outputSize int, truthTable map[int]bool) (g *Gate) {
	g = &Gate{
		ctx:        ctx,
		InputSize:  inputSize,
		OutputSize: outputSize,
		inputs:     make([]Transceiver, inputSize),
		outputs:    make([]Transceiver, outputSize),
		truthTable: truthTable,
		handlers:   make(map[HandlerSituation][]gateHandler),
		tick:       make(chan *sync.WaitGroup),
	}

	for _, situation := range HandlerSituations {
		g.handlers[situation] = make([]gateHandler, 0)
	}

	for i := 0; i < g.InputSize; i++ {
		g.inputs[i] = NewTransceiver()
	}

	for i := 0; i < g.OutputSize; i++ {
		g.outputs[i] = NewTransceiver()
	}

	go g.run()

	return
}

func (g *Gate) Outputs() (ts []Transmitter) {
	for _, output := range g.outputs {
		ts = append(ts, output)
	}

	return
}

func (g *Gate) Output(index int) (t Transmitter) {
	if g.OutputSize < index {
		return nil
	}

	return g.outputs[index]
}

func (g *Gate) Inputs() (rs []Receiver) {
	for _, input := range g.inputs {
		rs = append(rs, input)
	}

	return
}

func (g *Gate) Input(index int) (r Receiver) {
	if g.InputSize < index {
		return nil
	}

	return g.inputs[index]
}

func (g *Gate) run() {
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

		received := make([]bool, g.InputSize)
		state := make([]bool, g.InputSize)

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

func (g *Gate) getTruthTableIndex(state []bool) (index int) {
	for i, v := range state {
		index += boolMap[v] << i
	}

	return
}

func Connect(r Receiver, t Transmitter) {

}
