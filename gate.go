package logic_gate

import "sync"

var boolMap = map[bool]int{
	false: 0,
	true:  1,
}

type HandlerSituation int

const (
	AfterInput HandlerSituation = iota + 1
)

func (h HandlerSituation) String() string {
	switch h {
	case AfterInput:
		return "AfterInput"
	}
	return "unknown"
}

var HandlerSituations = []HandlerSituation{
	AfterInput,
}

type GateHandler func(g Gate, index int, input bool)

type Gate interface {
	Name() string
	InputSize() int
	Input(index int) Receiver // TODO: named input
	Inputs() []Receiver
	OutputSize() int
	Output(index int) Transmitter // TODO: named output
	Outputs() []Transmitter
	Tick(group *sync.WaitGroup)
	SetPreviousStatus(status bool)
	AddHandler(situation HandlerSituation, handler GateHandler)
}

func Connect(t Transmitter, r Receiver) {
	t.AppendReceiver(r)
}
