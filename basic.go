package logic_gate

import (
	"context"
)

func BasicGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
	}

	g = NewGate(ctx, 1, 1, truthTable)
	ConnectGateTicker(g)
	return
}

func PrintGate(ctx context.Context) (g *Gate) {
	g = BasicGate(ctx)
	g.handlers[AfterInput] = append(g.handlers[AfterInput], func(g *Gate, index int, input bool) {
		log.Println(g, index, input)
	})

	ConnectGateTicker(g)
	return g
}

func AndGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: true,
	}

	g = NewGate(ctx, 2, 1, truthTable)
	ConnectGateTicker(g)
	return g
}

func OrGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: true,
	}

	g = NewGate(ctx, 2, 1, truthTable)
	ConnectGateTicker(g)
	return g
}

func NotGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: true,
		1: false,
	}

	g = NewGate(ctx, 1, 1, truthTable)
	ConnectGateTicker(g)
	return g
}
