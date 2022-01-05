package logic_gate

import (
	"context"
)

func BasicGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
	}

	g = NewTruthTableGate(ctx, "BasicInput", 1, 1, truthTable)
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

	g = NewTruthTableGate(ctx, "AndGate", 2, 1, truthTable)
	return g
}

func OrGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: true,
	}

	g = NewTruthTableGate(ctx, "OrGate", 2, 1, truthTable)
	return g
}

func NotGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: true,
		1: false,
	}

	g = NewTruthTableGate(ctx, "NotGate", 1, 1, truthTable)
	return g
}

func XorGate(ctx context.Context) (g Gate) {
	truthTable := map[int]bool{
		0: false,
		1: true,
		2: true,
		3: false,
	}

	g = NewTruthTableGate(ctx, "XorGate", 2, 1, truthTable)
	return
}
