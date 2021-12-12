package logic_gate

import (
	"context"
	"reflect"
)

var boolMap = map[bool]int{
	false: 0,
	true:  1,
}

type Receiver interface {
	ConnectInput(Transmitter)
}

type Transmitter interface {
	ConnectOutput(Receiver)
}

type Gate struct {
	name           string
	ctx            context.Context
	InputSize      int
	OutputSize     int
	inputs         []chan bool
	outputs        []chan bool
	truthTable     map[int]bool
	state          []bool
	previousOutput bool
}

func NewGate(ctx context.Context, inputSize, outputSize int, truthTable map[int]bool) (g *Gate) {
	g = &Gate{
		ctx:        ctx,
		InputSize:  inputSize,
		OutputSize: outputSize,
		inputs:     make([]chan bool, 2),
		outputs:    make([]chan bool, 1),
		state:      make([]bool, 2),
		truthTable: truthTable,
	}

	for i := 0; i < g.InputSize; i++ {
		g.inputs[i] = make(chan bool, 0)
	}

	for i := 0; i < g.OutputSize; i++ {
		g.outputs[i] = make(chan bool, 0)
	}

	go g.run()

	return
}

func (g *Gate) run() {
	cases := make([]reflect.SelectCase, 0)

	for _, c := range g.inputs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}

	defer func() {
		for i := range g.outputs {
			close(g.outputs[i])
		}
	}()

	for {
		if err := g.ctx.Err(); err != nil {
			break
		}

		index, value, recvOk := reflect.Select(cases)
		if !recvOk {
			break
		}
		g.state[index] = value.Bool()

		next := g.truthTable[g.getTruthTableIndex()]
		//log.Println(index, value, g.name, g.state, g.previousOutput, "->", next)
		if next != g.previousOutput {
			g.outputs[0] <- next
			g.previousOutput = next
		}
	}
}

func (g *Gate) getTruthTableIndex() (index int) {
	for i, v := range g.state {
		index += boolMap[v] << i
	}

	return
}

//
//type ActionGate struct {
//	Gate
//	action func(inputs ...bool)
//}
//
//func PrintGate(ctx context.Context) ActionGate {
//	g := ActionGate{
//		Gate: Gate{
//			ctx: ctx,
//			InputSize: 1,
//			OutputSize: 1,
//			inputs: make([]chan bool, 1),
//			outputs: make([]chan bool, 1),
//			state: make([]bool, 1),
//			truthTable: map[int]bool {
//				0: false,
//				1: true,
//			},
//		},
//	}
//
//	for i:=0; i<g.InputSize; i++ {
//		g.inputs[i] = make(chan bool, 0)
//	}
//
//	for i:=0; i<g.OutputSize; i++ {
//		g.outputs[i] = make(chan bool, 0)
//	}
//
//	go g.run()
//
//	return g
//}
