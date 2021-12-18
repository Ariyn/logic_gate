package logic_gate

import (
	"context"
	"log"
)

func BasicGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
	}

	return NewGate(ctx, 2, 1, truthTable)
}

func PrintGate(ctx context.Context) (g *Gate) {
	g = BasicGate(ctx)
	g.handlers[AfterInput] = append(g.handlers[AfterInput], func(g *Gate, index int, input bool) {
		log.Println(g, index, input)
	})

	return g
}

func AndGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: true,
	}

	return NewGate(ctx, 2, 1, truthTable)
}

func OrGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: true,
	}

	return NewGate(ctx, 2, 1, truthTable)
}

func NotGate(ctx context.Context) (g *Gate) {
	truthTable := map[int]bool{
		0: true,
		1: false,
	}

	return NewGate(ctx, 1, 1, truthTable)
}
