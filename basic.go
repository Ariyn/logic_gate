package logic_gate

import (
	"context"
)

func BasicGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
	}

	g = NewTruthTableGate(ctx, 1, 1, truthTable)
	if GlobalEngine != nil {
		GlobalEngine.ConnectGateTicker(g)
	}
	return
}

//func PrintGate(ctx context.Context) (g Gate) {
//	g = BasicGate(ctx)
//	g.handlers[AfterInput] = append(g.handlers[AfterInput], func(g *TruthTableGate, index int, input bool) {
//		log.Println(g, index, input)
//	})
//
//	if globalEngine != nil {
//		globalEngine.ConnectGateTicker(g)
//	}
//	return g
//}

func AndGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: true,
	}

	g = NewTruthTableGate(ctx, 2, 1, truthTable)
	if GlobalEngine != nil {
		GlobalEngine.ConnectGateTicker(g)
	}

	return g
}

func OrGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: true,
	}

	g = NewTruthTableGate(ctx, 2, 1, truthTable)
	if GlobalEngine != nil {
		GlobalEngine.ConnectGateTicker(g)
	}

	return g
}

func NotGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: true,
		1: false,
	}

	g = NewTruthTableGate(ctx, 1, 1, truthTable)
	if GlobalEngine != nil {
		GlobalEngine.ConnectGateTicker(g)
	}

	return g
}

func XorGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: false,
	}

	g = NewTruthTableGate(ctx, 2, 1, truthTable)
	if GlobalEngine != nil {
		GlobalEngine.ConnectGateTicker(g)
	}

	return
}
