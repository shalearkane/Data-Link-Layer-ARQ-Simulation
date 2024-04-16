package transmitter

type Transmitter struct {
	In *chan int
}

func (p Transmitter) Transmit() {
	task := 1
	for {
		*p.In <- task
		task++
	}
}
