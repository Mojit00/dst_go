package recursion

import (
	"fmt"
	"math"
)

type EightQueue struct {
	QueueSize     int
	QueueLocation []int
	Count int
}

func NewQueue(size int) *EightQueue {
	queue := &EightQueue{
		QueueSize:     size,
		QueueLocation: make([]int, 8),
		Count:0,
	}
	return queue
}

func (q *EightQueue) Check(n int) {
	if n == q.QueueSize {
		q.print()
		return
	}
	for i := 0; i < q.QueueSize; i++ {
		q.QueueLocation[n] = i
		if q.judge(n) {
			q.Check(n + 1)
		}
	}
}

func (q *EightQueue) judge(cur int) bool {
	for i := 0; i < cur; i++ {
		if q.QueueLocation[i] == q.QueueLocation[cur] || math.Abs(float64(cur-i)) == math.Abs(float64(q.QueueLocation[cur]-q.QueueLocation[i])) {
			return false
		}
	}
	return true
}

func (q *EightQueue) print() {
	q.Count++
	fmt.Print("[")
	for _, v := range q.QueueLocation {
		fmt.Print(v, " ")
	}
	fmt.Println("]")
}
