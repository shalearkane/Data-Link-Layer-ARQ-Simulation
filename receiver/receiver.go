package receiver

import (
	"fmt"
)

type Receiver struct {
	In  *chan int
	Ack *chan bool
}

func (r Receiver) Receive() {
	count := 10
	for {
		message := <-*r.In
		println(fmt.Sprintf("Got a message %d", message))
		*r.Ack <- true
		count--
		if count < 0 {
			close(*r.In)
			close(*r.Ack)
			return
		}

	}
}
