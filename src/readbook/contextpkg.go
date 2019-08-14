package readbook

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context,wg *sync.WaitGroup,index int)error{
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello",index,"thread")
		case <-ctx.Done():
			fmt.Println("stop")
			return ctx.Err()
		}
	}
}

func CtxTest()  {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx,&wg,i)
	}

	time.Sleep(time.Second)

	cancle()

	wg.Wait()
}
