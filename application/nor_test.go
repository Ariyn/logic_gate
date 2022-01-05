package application

import (
	"context"
	"fmt"
	logic_gate "github.com/ariyn/logic-gate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNorGate(t *testing.T) {
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
				engine := logic_gate.NewEngine()

				g := NorGate(context.WithValue(context.TODO(), logic_gate.EngineKey, engine))

				g.Input(0).Push(tt[0])
				g.Input(1).Push(tt[1])

				engine.TickSync()
				engine.TickSync()

				assert.Equal(t, tt[2], g.Output(0).Pop())
			})
		}
	})
}
