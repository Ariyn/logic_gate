package logic_gate

import "sync"

var gates = make([]*Gate, 0)
var wg sync.WaitGroup

func ConnectGateTicker(g *Gate) {
	gates = append(gates, g)
}

func Tick() {
	for _, g := range gates {
		wg.Add(1)
		g.tick <- &wg
	}
	wg.Wait()

	wg = sync.WaitGroup{}
}
