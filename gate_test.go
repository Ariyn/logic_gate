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
		GlobalEngine.TickSync()
		GlobalEngine.TickSync()
		GlobalEngine.TickSync()
		GlobalEngine.TickSync()

		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())

		GlobalEngine.TickSync()
		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())

		GlobalEngine.TickSync()
		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())

		gate.Input(1).Push(false)
		GlobalEngine.TickSync()

		assert.True(t, gate.Output(0).Pop())
		assert.False(t, gate.Output(1).Pop())
	})
}

func TestGate_xor_gate(t *testing.T) {
	testCases := [][]bool{{
		false, false, false,
	}, {
		false, true, true,
	}, {
		true, false, true,
	}, {
		true, true, false,
	}}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
			xor := ComplexXorGate(context.TODO())

			xor.Input(0).Push(tt[0])
			xor.Input(1).Push(tt[1])
			for i := 0; i < 8; i++ {
				GlobalEngine.TickSync()
			}

			assert.Equal(t, tt[2], xor.Output(0).Pop())
		})
	}
}

func TestGate_FullAdder(t *testing.T) {
	// 0, 1, carry in, s, carry out
	testCases := [][]bool{{
		false, false, false, false, false,
	}, {
		false, false, true, true, false,
	}, {
		false, true, false, true, false,
	}, {
		false, true, true, false, true,
	}, {
		true, false, false, true, false,
	}, {
		true, false, true, false, true,
	}, {
		true, true, false, false, true,
	}, {
		true, true, true, true, true,
	}}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
			fullAdder := ComplexFullAdder(context.TODO())

			fullAdder.Input(0).Push(tt[0])
			fullAdder.Input(1).Push(tt[1])
			fullAdder.Input(2).Push(tt[2])
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()

			assert.Equal(t, tt[3], fullAdder.Output(0).Pop())
			assert.Equal(t, tt[4], fullAdder.Output(1).Pop())
		})
	}
}

func TestGate_2BitsFullAdder(t *testing.T) {
	// A0, A1, B0, B1, S0, S1, c2
	testCases := [][]bool{{
		false, false, false, false, false, false, false,
	}, {
		true, false, false, false, true, false, false,
	}, {
		false, true, false, false, false, true, false,
	}, {
		true, true, false, false, true, true, false,
	}, {
		false, false, true, false, true, false, false,
	}, {
		true, false, true, false, false, true, false,
	}, {
		false, true, true, false, true, true, false,
	}, {
		true, true, true, false, false, false, true,
	}, {
		false, false, false, true, false, true, false,
	}, {
		true, false, false, true, true, true, false,
	}, {
		false, true, false, true, false, false, true,
	}, {
		true, true, false, true, true, false, true,
	}, {
		false, false, true, true, true, true, false,
	}, {
		true, false, true, true, false, false, true,
	}, {
		false, true, true, true, true, false, true,
	}, {
		true, true, true, true, false, true, true,
	}}

	for index, tt := range testCases {
		t.Run(fmt.Sprintf("%d: when input case is (%v, %v), output is %v", index, tt[0], tt[1], tt[2]), func(t *testing.T) {
			fullAdder1 := ComplexFullAdder(context.TODO())
			fullAdder2 := ComplexFullAdder(context.TODO())

			Connect(fullAdder1.Output(1), fullAdder2.Input(2))

			fullAdder1.Input(0).Push(tt[0]) // A0
			fullAdder2.Input(0).Push(tt[1]) // A1
			fullAdder1.Input(1).Push(tt[2]) // B0
			fullAdder2.Input(1).Push(tt[3]) // B1

			GlobalEngine.TickSync()
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()
			GlobalEngine.TickSync()

			assert.Equal(t, tt[4], fullAdder1.Output(0).Pop(), "s0")
			assert.Equal(t, tt[5], fullAdder2.Output(0).Pop(), "s1")
			assert.Equal(t, tt[6], fullAdder2.Output(1).Pop(), "carry-out")
		})
	}
}

func TestGate_4BitsFullAdder(t *testing.T) {
	t.Run("0011 + 1100 = 1111", func(t *testing.T) {
		adder := Complex4BitsFullAdder(context.TODO())

		adder.Input(0).Push(true)
		adder.Input(1).Push(true)
		adder.Input(2).Push(false)
		adder.Input(3).Push(false)

		adder.Input(4).Push(false)
		adder.Input(5).Push(false)
		adder.Input(6).Push(true)
		adder.Input(7).Push(true)

		GlobalEngine.TickSync()
		GlobalEngine.TickSync()

		assert.True(t, adder.Output(0).Pop())
		assert.True(t, adder.Output(1).Pop())
		assert.True(t, adder.Output(2).Pop())
		assert.True(t, adder.Output(3).Pop())
		assert.False(t, adder.Output(4).Pop())
	})

	t.Run("1111 + 1100 = 11011", func(t *testing.T) {
		adder := Complex4BitsFullAdder(context.TODO())

		adder.Input(0).Push(true)
		adder.Input(1).Push(true)
		adder.Input(2).Push(true)
		adder.Input(3).Push(true)

		adder.Input(4).Push(false)
		adder.Input(5).Push(false)
		adder.Input(6).Push(true)
		adder.Input(7).Push(true)

		GlobalEngine.TickSync()
		GlobalEngine.TickSync()
		GlobalEngine.TickSync()

		assert.True(t, adder.Output(0).Pop())
		assert.True(t, adder.Output(1).Pop())
		assert.False(t, adder.Output(2).Pop())
		assert.True(t, adder.Output(3).Pop())
		assert.True(t, adder.Output(4).Pop())
	})
}

func TestGate_4BitsFullSubtractor(t *testing.T) {
	// 1100 = 12
	// 0011 = 3
	// 12 - 3 = 9 // 1001
	t.Run("1100 - 0011 = 1001", func(t *testing.T) {
		e := NewEngine()
		subtractor := Complex4BitsFullSubtractor(context.WithValue(context.Background(), EngineKey, e))

		subtractor.Input(0).Push(false)
		subtractor.Input(1).Push(false)
		subtractor.Input(2).Push(true)
		subtractor.Input(3).Push(true)

		subtractor.Input(4).Push(true)
		subtractor.Input(5).Push(true)
		subtractor.Input(6).Push(false)
		subtractor.Input(7).Push(false)

		e.TickSync()
		e.TickSync()
		e.TickSync()

		assert.True(t, subtractor.Output(0).Pop())
		assert.False(t, subtractor.Output(1).Pop())
		assert.False(t, subtractor.Output(2).Pop())
		assert.True(t, subtractor.Output(3).Pop())
	})

	t.Run("0011 - 0011 = 0000", func(t *testing.T) {
		e := NewEngine()
		subtractor := Complex4BitsFullSubtractor(context.WithValue(context.Background(), EngineKey, e))

		subtractor.Input(0).Push(true)
		subtractor.Input(1).Push(true)
		subtractor.Input(2).Push(false)
		subtractor.Input(3).Push(false)

		subtractor.Input(4).Push(true)
		subtractor.Input(5).Push(true)
		subtractor.Input(6).Push(false)
		subtractor.Input(7).Push(false)

		e.TickSync()
		e.TickSync()
		e.TickSync()
		e.TickSync()
		e.TickSync()

		assert.False(t, subtractor.Output(0).Pop())
		assert.False(t, subtractor.Output(1).Pop())
		assert.False(t, subtractor.Output(2).Pop())
		assert.False(t, subtractor.Output(3).Pop())
		assert.True(t, subtractor.Output(4).Pop())
	})

	t.Run("0000 - 0011 = underflow", func(t *testing.T) {
		e := NewEngine()
		subtractor := Complex4BitsFullSubtractor(context.WithValue(context.Background(), EngineKey, e))

		subtractor.Input(0).Push(false)
		subtractor.Input(1).Push(false)
		subtractor.Input(2).Push(false)
		subtractor.Input(3).Push(false)

		subtractor.Input(4).Push(true)
		subtractor.Input(5).Push(true)
		subtractor.Input(6).Push(false)
		subtractor.Input(7).Push(false)

		e.TickSync()
		e.TickSync()
		e.TickSync()

		assert.True(t, subtractor.Output(0).Pop())
		assert.False(t, subtractor.Output(1).Pop())
		assert.True(t, subtractor.Output(2).Pop())
		assert.True(t, subtractor.Output(3).Pop())
		assert.False(t, subtractor.Output(4).Pop())
	})
}
