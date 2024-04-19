package receiver

import (
	"context"
	"fmt"
)

type Receiver struct {
	In *chan int
}

func (c Receiver) Receive(ctx context.Context) {
	for {
		select {
		case job := <-*c.In:
			{
				println(fmt.Sprintf("Got a job %d", job))
			}
		case <-ctx.Done():
			return
		}
	}
}
