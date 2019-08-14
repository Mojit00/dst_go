package readbook

import (
	"fmt"
	"testing"
	"time"
)

func TestMethod1(t *testing.T) {
	Method1()
}

func TestMethod2(t *testing.T) {
	Method2()
}

func TestMethod3(t *testing.T) {
	Method3()
}

func TestMethod4(t *testing.T) {
	Method4()
}

func TestConsumer(t *testing.T) {
	ch:= make(chan int, 64)

	go Producer(3,ch)
	go Producer(5,ch)
	go Consumer(ch)
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("app is ready to stop")
	}
}