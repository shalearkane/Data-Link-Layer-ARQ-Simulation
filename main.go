package main

import (
	"channels/receiver"
	"channels/transmitter"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	in := make(chan int, 1)
	ack := make(chan bool, 1)

	t := transmitter.Transmitter{In: &in, Ack: &ack}
	r := receiver.Receiver{In: &in, Ack: &ack}
	go t.Transmit()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go r.Receive(ctx)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancelFunc()
}
