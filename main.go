package main

import (
	"channels/receiver"
	"channels/transmitter"
	"context"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func main() {
	const nConsumers = 10
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := make(chan int, 1)

	t := transmitter.Transmitter{In: &in}
	r := receiver.Receiver{In: &in, Jobs: make(chan int, nConsumers)}
	go t.Transmit()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go r.Receive(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(nConsumers)
	for i := 0; i < nConsumers; i++ {
		go r.Work(wg)
	}
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancelFunc()
	wg.Wait()
}
