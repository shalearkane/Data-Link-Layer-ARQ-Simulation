package transmitter

import (
	"fmt"
)

type Transmitter struct {
	In  *chan []byte
	Ack *chan bool
}

func (t Transmitter) Transmit() {
	*t.Ack <- true
	message := 1

	for {
		ack := <-*t.Ack
		if ack {
			fmt.Println(ack)
			data := make([]byte, 64)
			crcAddedData := AddCRC(&data)
			*t.In <- AddNoise(crcAddedData)
			message++
		} else {
			message--
			*t.In <- make([]byte, 64)
		}

	}
}
