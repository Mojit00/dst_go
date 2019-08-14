package recursion

import (
	"fmt"
	"testing"
)

func TestEightQueue_Check(t *testing.T) {
	queue := NewQueue(8)
	queue.Check(0)
	fmt.Println("一共有解法:",queue.Count)
}
