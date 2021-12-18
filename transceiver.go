package logic_gate

type Receiver interface{}
type Transmitter interface{}

type Transceiver struct {
	status   bool
	input    chan bool
	outputs  []chan bool
	outputs2 []*Transceiver
}

func NewTransceiver() Transceiver {
	return Transceiver{
		input:    make(chan bool, 1),
		outputs:  make([]chan bool, 0),
		outputs2: make([]*Transceiver, 0),
	}
}

func (t *Transceiver) Close() {
	for _, o := range t.outputs {
		close(o)
	}
}

func (t *Transceiver) Push(signal bool) {
	t.status = signal
}

func (t *Transceiver) Pop() (signal bool) {
	return t.status
}

func (t *Transceiver) receive() (received bool) {
	select {
	case signal := <-t.input:
		t.status = signal
		received = true
	default:
	}
	return
}

func (t *Transceiver) transmit(status bool) (transmitted int) {
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
