package logic_gate

import "sync"

var GlobalEngine *Engine

func init() {
	GlobalEngine = NewEngine()
}

type Engine struct {
	gates []Gate
	wg    sync.WaitGroup
}

func NewEngine() (e *Engine) {
	return &Engine{
		gates: make([]Gate, 0),
	}
}

func (e *Engine) ConnectGateTicker(g Gate) {
	e.gates = append(e.gates, g)
}

// TODO: record every gate's status
func (e *Engine) TickSync() {
	e.wg = sync.WaitGroup{}

	for _, g := range e.gates {
		e.wg.Add(1)
		g.Tick(&e.wg)
	}

	e.wg.Wait()
}
