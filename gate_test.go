package logic_gate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestPrintGate(t *testing.T) {
	t.Run("basic gate", func(t *testing.T) {
		g := BasicGate(context.Background())

		g.inputs[0] <- true
		assert.True(t, <-g.outputs[0])

		g.inputs[0] <- false
		assert.False(t, <-g.outputs[0])
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
		g := AndGate(context.Background())

		for _, tt := range testCases {
			g.inputs[0] <- false
			<-g.outputs[0]
			g.inputs[1] <- false
			<-g.outputs[0]

			g.inputs[0] <- tt[0]
			<-g.outputs[0]

			g.inputs[1] <- tt[1]

			assert.Equal(t, tt[2], <-g.outputs[0])
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
		g := OrGate(context.Background())

		for _, tt := range testCases {
			g.inputs[0] <- false
			<-g.outputs[0]
			g.inputs[1] <- false
			<-g.outputs[0]

			g.inputs[0] <- tt[0]
			<-g.outputs[0]

			g.inputs[1] <- tt[1]

			assert.Equal(t, tt[2], <-g.outputs[0])
		}
	})

	t.Run("not gate", func(t *testing.T) {
		g := NotGate(context.Background())

		g.inputs[0] <- true
		assert.False(t, <-g.outputs[0])

		g.inputs[0] <- false
		assert.True(t, <-g.outputs[0])
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
		g := NorGate(context.Background())

		for _, tt := range testCases {
			g.inputs[0] <- false
			<-g.outputs[0]
			g.inputs[1] <- false
			<-g.outputs[0]

			g.inputs[0] <- tt[0]
			<-g.outputs[0]

			g.inputs[1] <- tt[1]

			assert.Equal(t, tt[2], <-g.outputs[0])
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
		g := NandGate(context.Background())

		for _, tt := range testCases {
			g.inputs[0] <- false
			<-g.outputs[0]
			g.inputs[1] <- false
			<-g.outputs[0]

			g.inputs[0] <- tt[0]
			<-g.outputs[0]

			g.inputs[1] <- tt[1]

			assert.Equal(t, tt[2], <-g.outputs[0])
		}
	})
}

func TestGate_connection(t *testing.T) {
	t.Run("connect 2 gates", func(t *testing.T) {
		g1 := BasicGate(context.Background())
		g2 := NotGate(context.Background())

		g2.inputs[0] = g1.outputs[0]

		g1.inputs[0] <- true
		assert.False(t, <-g2.outputs[0])

		g1.inputs[0] <- false
		assert.True(t, <-g2.outputs[0])
	})
}

func TestGate_flipflop(t *testing.T) {
	t.Run("nor ratch", func(t *testing.T) {
		rGate := NorGate(context.Background())
		rGate.gates[0].name = "r gate"
		sGate := NorGate(context.Background())
		sGate.gates[0].name = "s gate"

		sGate.inputs[1] = rGate.outputs[0]
		rGate.inputs[1] = sGate.outputs[0]

		rGate.gates[0].state[1] = true
		rGate.gates[0].previousOutput = false
		rGate.gates[1].previousOutput = true
		sGate.gates[0].previousOutput = true

		log.Println("r on")
		rGate.inputs[0] <- true
		time.Sleep(time.Millisecond * 300)

		log.Println("r off")
		rGate.inputs[0] <- false
		time.Sleep(time.Millisecond * 300)

		log.Println("s on")
		sGate.inputs[0] <- true
		time.Sleep(time.Millisecond * 300)

		log.Println("s off")
		sGate.inputs[0] <- false
	})
}
