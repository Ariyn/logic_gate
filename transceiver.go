package logic_gate

type Receiver interface {
	Input() chan bool
	Push(status bool)
	Receive() bool // TODO: should both Receive and Push exist?
	Status() bool
	SetCurrentStatus(status bool)
}

type Transmitter interface {
	Pop() bool
	Transmit(bool) int // TODO: should both Transmit and Pop exist?
	AppendReceiver(r Receiver)
	Close()
}

var _ Receiver = (*Transceiver)(nil)
var _ Transmitter = (*Transceiver)(nil)

type Transceiver struct {
	status  bool
	input   chan bool
	outputs []chan bool
}

func NewTransceiver() *Transceiver {
	return &Transceiver{
		input:   make(chan bool, 1),
		outputs: make([]chan bool, 0),
	}
}

func (t *Transceiver) Close() {
	for _, o := range t.outputs {
		close(o)
	}
}

func (t *Transceiver) Status() bool {
	return t.status
}

func (t *Transceiver) Push(signal bool) {
	t.status = signal
}

func (t *Transceiver) Pop() (signal bool) {
	return t.status
}

func (t *Transceiver) Receive() (received bool) {
	select {
	case signal := <-t.input:
		t.status = signal
		received = true
	default:
	}
	return
}

func (t *Transceiver) Transmit(status bool) (transmitted int) {
	t.status = status

	transmitted = 0
	for _, output := range t.outputs {
		select {
		case output <- t.status:
			transmitted += 1
		default:
		}
	}

	return
}

func (t *Transceiver) SetCurrentStatus(status bool) {
	t.status = status
}

func (t *Transceiver) AppendReceiver(r Receiver) {
	t.outputs = append(t.outputs, r.Input())
}

func (t *Transceiver) Input() chan bool {
	return t.input
}
