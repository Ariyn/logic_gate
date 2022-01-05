package application

import (
	"context"
	logicGate "github.com/ariyn/logic-gate"
)

func NewMux(ctx context.Context) (g logicGate.Gate) {
	s := []logicGate.Gate{
		logicGate.BasicGate(ctx),
		logicGate.BasicGate(ctx),
	}

	sNot := []logicGate.Gate{
		logicGate.NotGate(ctx),
		logicGate.NotGate(ctx),
		logicGate.NotGate(ctx),
		logicGate.NotGate(ctx),
	}

	logicGate.Connect(s[0].Output(0), sNot[0].Input(0))
	logicGate.Connect(s[1].Output(0), sNot[1].Input(0))

	input := []logicGate.Gate{
		logicGate.BasicGate(ctx),
		logicGate.BasicGate(ctx),
		logicGate.BasicGate(ctx),
		logicGate.BasicGate(ctx),
	}

	// TODO: nbit and gate가 필요하다...
	sAnd := []logicGate.Gate{
		NBitsAnd(ctx, 3),
		NBitsAnd(ctx, 3),
		NBitsAnd(ctx, 3),
		NBitsAnd(ctx, 3),
	}

	logicGate.Connect(input[0].Output(0), sAnd[0].Input(2))
	logicGate.Connect(input[1].Output(0), sAnd[1].Input(2))
	logicGate.Connect(input[2].Output(0), sAnd[2].Input(2))
	logicGate.Connect(input[3].Output(0), sAnd[3].Input(2))

	logicGate.Connect(sNot[0].Output(0), sAnd[0].Input(0))
	logicGate.Connect(sNot[1].Output(0), sAnd[0].Input(1))

	logicGate.Connect(s[0].Output(0), sAnd[1].Input(0))
	logicGate.Connect(sNot[1].Output(0), sAnd[1].Input(1))

	logicGate.Connect(sNot[0].Output(0), sAnd[2].Input(0))
	logicGate.Connect(s[1].Output(0), sAnd[2].Input(1))

	logicGate.Connect(s[0].Output(0), sAnd[3].Input(0))
	logicGate.Connect(s[1].Output(0), sAnd[3].Input(1))

	or := NBitsOr(ctx, 4)
	logicGate.Connect(sAnd[0].Output(0), or.Input(0))
	logicGate.Connect(sAnd[1].Output(0), or.Input(1))
	logicGate.Connect(sAnd[2].Output(0), or.Input(2))
	logicGate.Connect(sAnd[3].Output(0), or.Input(3))

	return logicGate.NewComplexGate(
		ctx,
		"4bit mux",
		[]logicGate.Receiver{s[0].Input(0), s[1].Input(0), input[0].Input(0), input[1].Input(0), input[2].Input(0), input[3].Input(0)},
		[]logicGate.Transmitter{or.Output(0)},
		[]logicGate.Gate{s[0], s[1], sNot[0], sNot[1], input[0], input[1], input[2], input[3], sAnd[0], sAnd[1], sAnd[2], sAnd[3]},
	)
}
