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

var HandlerSituations = []HandlerSituation{
	AfterInput,
}

type Gate interface {
	InputSize() int
	Input(index int) Receiver
	Inputs() []Receiver
	OutputSize() int
	Output(index int) Transmitter
	Outputs() []Transmitter
	Tick(group *sync.WaitGroup)
	SetPreviousStatus(status bool)
}

type gateHandler func(g *Gate, index int, input bool)

func Connect(t Transmitter, r Receiver) {
	t.AppendReceiver(r)
}