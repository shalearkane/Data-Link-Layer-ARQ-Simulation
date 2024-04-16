package receiver

import (
	"context"
	"fmt"
	"sync"
)

type Receiver struct {
	In   *chan int
	Jobs chan int
}

func (c Receiver) Work(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range c.Jobs {
		fmt.Printf("%dth job finished\n", job)
	}
}
func (c Receiver) Receive(ctx context.Context) {
	for {
		select {
		case job := <-*c.In:
			c.Jobs <- job
		case <-ctx.Done():
			close(c.Jobs)
			return
		}
	}
}
