package application

import (
	"context"
	"fmt"
	logicGate "github.com/ariyn/logic-gate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNBitsOr(t *testing.T) {
	t.Run("3bit or gate test", func(t *testing.T) {
		testCase := make([][4]bool, 8)
		for i := 0; i < 8; i++ {
			testCase[i] = [4]bool{}

			for j := 0; j < 3; j++ {
				testCase[i][j] = (i >> j & 1) == 1
			}
			testCase[i][3] = true
		}

		testCase[0][3] = false

		for index, tt := range testCase {
			t.Run(fmt.Sprintf("%d %v", index, tt[3]), func(t *testing.T) {
				newEngine := logicGate.NewEngine()
				ctx := context.WithValue(context.TODO(), logicGate.EngineKey, newEngine)

				orGate := NBitsOr(ctx, 3)
				orGate.Input(0).Push(tt[0])
				orGate.Input(1).Push(tt[1])
				orGate.Input(2).Push(tt[2])

				newEngine.TickSync()

				assert.Equal(t, tt[3], orGate.Output(0).Pop())
			})
		}
	})
}
