package main

import (
	"channels/receiver"
	"channels/transmitter"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	in := make(chan []byte, 1)
	ack := make(chan bool, 1)

	t := transmitter.Transmitter{In: &in, Ack: &ack}
	r := receiver.Receiver{In: &in, Ack: &ack}
	go t.Transmit()
	go r.Receive()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
}
