package logic_gate

import (
	"context"
	"sync"
)

var _ Gate = (*TruthTableGate)(nil)

type TruthTableGate struct {
	name            string
	ctx             context.Context
	receiverSize    int
	receivers       []Receiver
	transmitterSize int
	transmitters    []Transmitter
	truthTable      map[int]bool
	previousOutput  bool
	handlers        map[HandlerSituation][]GateHandler
	tick            chan *sync.WaitGroup
	isRunning       bool
}

func NewTruthTableGate(ctx context.Context, name string, inputSize, outputSize int, truthTable map[int]bool) (g Gate) {
	tg := &TruthTableGate{
		ctx:             ctx,
		name:            name,
		receiverSize:    inputSize,
		transmitterSize: outputSize,
		receivers:       make([]Receiver, inputSize),
		transmitters:    make([]Transmitter, outputSize),
		truthTable:      truthTable,
		handlers:        make(map[HandlerSituation][]GateHandler),
		tick:            make(chan *sync.WaitGroup),
	}

	for _, situation := range HandlerSituations {
		tg.handlers[situation] = make([]GateHandler, 0)
	}

	for i := 0; i < tg.receiverSize; i++ {
		tg.receivers[i] = NewTransceiver()
	}

	for i := 0; i < tg.transmitterSize; i++ {
		tg.transmitters[i] = NewTransceiver()
	}

	go tg.run()

	return tg
}

func (g *TruthTableGate) Name() string {
	return g.name
}

func (g *TruthTableGate) InputSize() (size int) {
	return g.receiverSize
}

func (g *TruthTableGate) Input(index int) (r Receiver) {
	if g.receiverSize < index {
		return nil
	}

	return g.receivers[index]
}

func (g *TruthTableGate) Inputs() (rs []Receiver) {
	for _, r := range g.receivers {
		rs = append(rs, r)
	}

	return
}

func (g *TruthTableGate) OutputSize() int {
	return g.transmitterSize
}

func (g *TruthTableGate) Output(index int) (t Transmitter) {
	if g.transmitterSize < index {
		return nil
	}

	return g.transmitters[index]
}

func (g *TruthTableGate) Outputs() (ts []Transmitter) {
	for _, r := range g.transmitters {
		ts = append(ts, r)
	}

	return
}

func (g *TruthTableGate) Tick(wg *sync.WaitGroup) {
	g.tick <- wg
}

func (g *TruthTableGate) run() {
	defer func() {
		for _, o := range g.transmitters {
			o.Close()
		}
	}()

	g.isRunning = true

	for g.isRunning {
		select {
		case <-g.ctx.Done():
			g.isRunning = false
			break
		case wg := <-g.tick:
			received := make([]bool, g.receiverSize)
			state := make([]bool, g.receiverSize)

			for i := range g.receivers {
				received[i] = g.receivers[i].Receive()
				state[i] = g.receivers[i].Status()
			}

			if len(g.handlers[AfterInput]) != 0 {
				for _, f := range g.handlers[AfterInput] {
					for i := range state {
						f(g, i, state[i])
					}
				}
			}

			// TODO: TruthTableGate should return more than 2 outputs
			next := g.truthTable[g.getTruthTableIndex(state)]
			if next != g.previousOutput {
				g.transmitters[0].Transmit(next)
				g.previousOutput = next
			}

			wg.Done()
		}
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

func (g *TruthTableGate) AddHandler(situation HandlerSituation, handler GateHandler) {
	if _, ok := g.handlers[situation]; ok {
		g.handlers[situation] = append(g.handlers[situation], handler)
	}
}
