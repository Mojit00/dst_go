package readbook

import (
	"fmt"
	"sync"
)

func Method1() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello,world")
		mu.Unlock()
	}()

	mu.Lock()
}

func Method2() {
	done := make(chan int)

	go func() {
		fmt.Println("hello world")
		done <- 1
	}()

	<-done
}

func Method3() {
	done := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("hello world")
			done <- 1
		}()
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func Method4() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello,world")
			wg.Done()
		}()
	}

	wg.Wait()
}

func Producer(factor int, in chan<- int) {
	for i := 0; i < 1000; i++ {
		in <- i * factor
		fmt.Println("producer value:",i * factor)
	}
}

func Consumer(out <-chan int){
	for value := range out {
		fmt.Println("consumer consume value:",value)
	}
}

