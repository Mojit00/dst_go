package LinkedList

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type CircleList struct {
	First *Node
	Size  int
}

func (l *CircleList) AddNode(num int) {
	var curBoy *Node
	if num < 1 {
		fmt.Print("num must be over 1")
		return
	}
	l.Size = num
	for i := 1; i <= num; i++ {
		boy := &Node{Value: i}
		if l.First == nil {
			l.First = boy
			boy.Next = l.First
			curBoy = l.First
		} else {
			curBoy.Next = boy
			boy.Next = l.First
			curBoy = boy
		}
	}
}

func (l *CircleList) Show() {
	cur := l.First
	if cur == nil {
		return
	}
	for {
		fmt.Println(cur.Value)
		if cur.Next == l.First {
			break
		}
		cur = cur.Next
	}
}

func (l *CircleList) CountBoy(step int) {
	if step < 0 || step > l.Size {
		fmt.Println("step is too long")
		return
	}
	if l.First == nil {
		return
	}
	helper := l.First

	//找到目标位置的前一个
	for true {
		if helper.Next == l.First {
			break
		}
		helper = helper.Next
	}
	//现在去位移
	for true {
		if helper == l.First {
			break
		}

		for i := 0; i < step-1; i++ {
			helper = helper.Next
			l.First = l.First.Next
		}

		fmt.Println("出队", l.First.Value)
		l.First = l.First.Next;
		helper.Next = l.First;
	}
	fmt.Println("最后", helper.Value)
}
