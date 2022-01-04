package application

import (
	"context"
	"fmt"
	logicGate "github.com/ariyn/logic-gate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGate_HalfAdder(t *testing.T) {
	testCases := [][]bool{{
		false, false, false, false,
	}, {
		false, true, true, false,
	}, {
		true, false, true, false,
	}, {
		true, true, false, true,
	}}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("when input case is (%v, %v), output is %v", tt[0], tt[1], tt[2]), func(t *testing.T) {
			hfAdder := HalfAdder(context.TODO())
			hfAdder.Input(0).Push(tt[0])
			hfAdder.Input(1).Push(tt[1])
			logicGate.GlobalEngine.TickSync()
			logicGate.GlobalEngine.TickSync()

			assert.Equal(t, tt[2], hfAdder.Output(0).Pop())
			assert.Equal(t, tt[3], hfAdder.Output(1).Pop())
		})
	}
}
