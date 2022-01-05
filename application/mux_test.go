package application

import (
	"context"
	"fmt"
	logicGate "github.com/ariyn/logic-gate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMuxBasicMuxing(t *testing.T) {
	testCases := []struct {
		selector int
		status   []bool
	}{{
		selector: 0,
		status:   []bool{true, false, false, false},
	}, {
		selector: 1,
		status:   []bool{false, true, false, false},
	}, {
		selector: 2,
		status:   []bool{false, false, true, false},
	}, {
		selector: 3,
		status:   []bool{false, false, false, true},
	}}

	for index, tt := range testCases {
		t.Run(fmt.Sprintf("%d test, select %d", index, tt.selector), func(t *testing.T) {
			engine := logicGate.NewEngine()
			ctx := context.WithValue(context.TODO(), logicGate.EngineKey, engine)
			mux := NewMux(ctx)

			mux.Input(2).Push(tt.status[0])
			mux.Input(3).Push(tt.status[1])
			mux.Input(4).Push(tt.status[2])
			mux.Input(5).Push(tt.status[3])

			for i := 0; i < 4; i++ {
				mux.Input(0).Push(i&1 == 1)
				mux.Input(1).Push(i>>1&1 == 1)

				engine.TickSync()
				engine.TickSync()
				engine.TickSync()

				if i == tt.selector {
					assert.True(t, mux.Output(0).Pop())
				} else {
					assert.False(t, mux.Output(0).Pop())
				}
			}
		})
	}
}
