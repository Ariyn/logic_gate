package logic_gate

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGate(t *testing.T) {
	t.Run("basic gate", func(t *testing.T) {
		g := BasicGate(context.Background())

		t.Run("true를 넣었을 때, true가 잘 나오는지", func(t *testing.T) {
			g.inputs[0].Push(true)
			Tick()
			assert.True(t, g.outputs[0].Pop())
		})

		t.Run("false를 넣었을 때, false가 잘 나오는지", func(t *testing.T) {
			g.inputs[0].Push(false)
			Tick()
			assert.False(t, g.outputs[0].Pop())
		})
	})

	//t.Run("print gate", func(t *testing.T) {
	//	g := PrintGate(context.Background())
	//
	//	g.inputs[0] <- true
	//	assert.True(t, <-g.outputs[0])
	//
	//	g.inputs[0] <- false
	//	assert.False(t, <-g.outputs[0])
	//})

	t.Run("and gate", func(t *testing.T) {
		testCases := [][]bool{{
			false, false, false,
		}, {
			false, true, false,
		}, {
			true, false, false,
		}, {
			true, true, true,
		}}

		for _, tt := range testCases {
			t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
				g := AndGate(context.Background())

				g.inputs[0].Push(tt[0])
				g.inputs[1].Push(tt[1])
				Tick()

				assert.Equal(t, tt[2], g.outputs[0].Pop())
			})
		}
	})

	t.Run("or gate", func(t *testing.T) {
		testCases := [][]bool{{
			false, false, false,
		}, {
			false, true, true,
		}, {
			true, false, true,
		}, {
			true, true, true,
		}}

		for _, tt := range testCases {
			t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
				g := OrGate(context.Background())

				g.inputs[0].Push(tt[0])
				g.inputs[1].Push(tt[1])
				Tick()

				assert.Equal(t, tt[2], g.outputs[0].Pop())
			})
		}
	})

	t.Run("not gate", func(t *testing.T) {
		g := NotGate(context.Background())

		t.Run("true가 입력되었을 때, false가 리턴됨", func(t *testing.T) {
			g.inputs[0].Push(true)
			Tick()

			assert.False(t, g.outputs[0].Pop())
		})

		t.Run("false가 입력 되었을 때, true가 반환됨", func(t *testing.T) {
			g.inputs[0].Push(false)
			Tick()

			assert.True(t, g.outputs[0].Pop())
		})
	})

	t.Run("nor gate", func(t *testing.T) {
		testCases := [][]bool{{
			false, false, true,
		}, {
			false, true, false,
		}, {
			true, false, false,
		}, {
			true, true, false,
		}}

		for _, tt := range testCases {
			t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
				g := NorGate(context.Background())

				g.inputs[0].Push(tt[0])
				g.inputs[1].Push(tt[1])
				Tick()
				Tick()

				assert.Equal(t, tt[2], g.outputs[0].Pop())
			})
		}
	})

	t.Run("nand gate", func(t *testing.T) {
		testCases := [][]bool{{
			false, false, true,
		}, {
			false, true, true,
		}, {
			true, false, true,
		}, {
			true, true, false,
		}}

		for _, tt := range testCases {
			t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
				g := NandGate(context.Background())

				g.inputs[0].Push(tt[0])
				g.inputs[1].Push(tt[1])
				Tick()
				Tick()

				assert.Equal(t, tt[2], g.outputs[0].Pop())
			})
		}
	})
}

func TestGate_flipflop(t *testing.T) {
	t.Run("nor ratch", func(t *testing.T) {
		rGate := NorGate(context.Background())
		rGate.gates[0].name = "r gate"
		sGate := NorGate(context.Background())
		sGate.gates[0].name = "s gate"

		rGate.outputs[0].outputs = append(rGate.outputs[0].outputs, sGate.inputs[1].input)
		sGate.outputs[0].outputs = append(sGate.outputs[0].outputs, rGate.inputs[1].input)

		rGate.gates[0].inputs[1].status = true
		rGate.gates[0].previousOutput = false
		rGate.gates[1].previousOutput = true
		sGate.gates[0].previousOutput = true

		log.Println("r on")
		rGate.inputs[0].Push(true)
		Tick()
		log.Println("   ", rGate.outputs[0].Pop(), sGate.outputs[0].Pop())

		log.Println("r off")
		rGate.inputs[0].Push(false)
		Tick()
		log.Println("   ", rGate.outputs[0].Pop(), sGate.outputs[0].Pop())

		log.Println("s on")
		sGate.inputs[0].Push(true)
		Tick()
		Tick()
		log.Println("   ", rGate.outputs[0].Pop(), sGate.outputs[0].Pop())

		log.Println("s off")
		sGate.inputs[0].Push(false)
		Tick()
		log.Println("   ", rGate.outputs[0].Pop(), sGate.outputs[0].Pop())
	})
}

//func TestGate_Receiver(t *testing.T) {
//	g := BasicGate(context.TODO())
//	printer := PrintGate(context.TODO())
//	Connect(printer.Input(0), g.Output(0))
//}
