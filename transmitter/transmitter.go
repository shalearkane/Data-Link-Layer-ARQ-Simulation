package transmitter

import (
	"fmt"
)

type Transmitter struct {
	In  *chan int
	Ack *chan bool
}

func (t Transmitter) Transmit() {
	*t.Ack <- true
	message := 1

	for {
		ack := <-*t.Ack
		if ack {
			fmt.Println(ack)
			*t.In <- message
			message++
		} else {
			message--
			*t.In <- message
		}

	}
}
