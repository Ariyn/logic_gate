package logic_gate

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGate(t *testing.T) {
	t.Run("basic gate", func(t *testing.T) {
		g := BasicGate(context.Background())

		t.Run("true를 넣었을 때, true가 잘 나오는지", func(t *testing.T) {
			g.Input(0).Push(true)
			GlobalEngine.TickSync()
			assert.True(t, g.Output(0).Pop())
		})

		t.Run("false를 넣었을 때, false가 잘 나오는지", func(t *testing.T) {
			g.Input(0).Push(false)
			GlobalEngine.TickSync()
			assert.False(t, g.Output(0).Pop())
		})
	})

	//t.Run("print gate", func(t *testing.T) {
	//	g := PrintGate(context.Background())
	//
	//	g.inputs[0] <- true
	//	assert.True(t, <-g.Output(0))
	//
	//	g.inputs[0] <- false
	//	assert.False(t, <-g.Output(0))
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

				g.Input(0).Push(tt[0])
				g.Input(1).Push(tt[1])
				GlobalEngine.TickSync()

				assert.Equal(t, tt[2], g.Output(0).Pop())
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

				g.Input(0).Push(tt[0])
				g.Input(1).Push(tt[1])
				GlobalEngine.TickSync()

				assert.Equal(t, tt[2], g.Output(0).Pop())
			})
		}
	})

	t.Run("not gate", func(t *testing.T) {
		g := NotGate(context.Background())

		t.Run("true가 입력되었을 때, false가 리턴됨", func(t *testing.T) {
			g.Input(0).Push(true)
			GlobalEngine.TickSync()

			assert.False(t, g.Output(0).Pop())
		})

		t.Run("false가 입력 되었을 때, true가 반환됨", func(t *testing.T) {
			g.Input(0).Push(false)
			GlobalEngine.TickSync()

			assert.True(t, g.Output(0).Pop())
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

				g.Input(0).Push(tt[0])
				g.Input(1).Push(tt[1])

				// need 2 ticks for or gate and not gate both work
				GlobalEngine.TickSync()
				GlobalEngine.TickSync()

				assert.Equal(t, tt[2], g.Output(0).Pop())
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

				g.Input(0).Push(tt[0])
				g.Input(1).Push(tt[1])

				// need 2 ticks for and gate and not gate both work
				GlobalEngine.TickSync()
				GlobalEngine.TickSync()

				assert.Equal(t, tt[2], g.Output(0).Pop())
			})
		}
	})
}

func TestGate_flipflop(t *testing.T) {
	t.Run("nor ratch", func(t *testing.T) {
		gate := FlipFlopSR(context.TODO())

		gate.Input(0).Push(true)
		GlobalEngine.TickSync()

		assert.False(t, gate.Output(0).Pop())
		assert.True(t, gate.Output(1).Pop())

		gate.Input(0).Push(false)
		GlobalEngine.TickSync()

		assert.False(t, gate.Output(0).Pop())
		assert.True(t, gate.Output(1).Pop())

		// FIX: sometimes, sGate does not work well in 2 ticks.
		// but it should work in 2 ticks.
		gate.Input(1).Push(true)
		globalEngine.TickSync()
		globalEngine.TickSync()
		globalEngine.TickSync()

		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())

		gate.Input(1).Push(false)
		globalEngine.TickSync()

		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())
	})
}

//func TestGate_Receiver(t *testing.T) {
//	g := BasicGate(context.TODO())
//	printer := PrintGate(context.TODO())
//	Connect(printer.Input(0), g.Output(0))
//}
