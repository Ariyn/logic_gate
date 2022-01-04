package logic_gate

import (
	"fmt"
	"strings"
	"sync"
)

const EngineKey = "_ENGINE"

var GlobalEngine *Engine

func init() {
	GlobalEngine = NewEngine()
}

type debugStatus struct {
	gate   Gate
	index  int
	status bool
}

// TODO: Rename Engine to Ticker
type Engine struct {
	gates         []Gate
	wg            sync.WaitGroup
	tick          int
	debugStatuses map[int][]debugStatus
}

func NewEngine() (e *Engine) {
	return &Engine{
		gates:         make([]Gate, 0),
		debugStatuses: make(map[int][]debugStatus),
	}
}

func (e *Engine) ConnectGateTicker(g Gate) (index int) {
	e.gates = append(e.gates, g)

	return len(e.gates)
}

func (e *Engine) SetDebuggingMode() {
	for i := range e.gates {
		e.gates[i].AddHandler(AfterInput, func(g Gate, index int, input bool) {
			e.writeStatus(AfterInput, g, index, input)
		})
	}
}

func (e *Engine) writeStatus(situation HandlerSituation, g Gate, index int, status bool) {
	//log.Println(e.tick, situation, g.Name(), index, status)
	e.debugStatuses[e.tick] = append(e.debugStatuses[e.tick], debugStatus{
		gate:   g,
		index:  index,
		status: status,
	})
}

func (e *Engine) WriteStatus() {
	datas := []string{fmt.Sprintf("tick %d", e.tick)}

	for _, s := range e.debugStatuses[e.tick] {
		datas = append(datas, fmt.Sprintf("\t%s, %d, %v", s.gate.Name(), s.index, s.status))
	}

	fmt.Println(strings.Join(datas, "\n") + "\n")
}

// TODO: record every gate's status
func (e *Engine) TickSync() {
	e.debugStatuses[e.tick] = make([]debugStatus, 0)
	e.wg = sync.WaitGroup{}

	for _, g := range e.gates {
		e.wg.Add(1)
		g.Tick(&e.wg)
	}

	e.wg.Wait()
	e.WriteStatus()

	e.tick += 1
}
