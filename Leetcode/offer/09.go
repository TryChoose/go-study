package main

import "fmt"

//剑指offer 用两个栈实现队列

type Queue struct {
	i, o []int
}

func NewQueue() *Queue {
	return &Queue{}
}

// 添加元素
func (q *Queue) Add(v int) {
	q.i = append(q.i, v)
}

// 删除元素
func (q *Queue) Del() int {
	if len(q.o) == 0 {
		if len(q.i) == 0 {
			return -1
		}
		for len(q.i) > 0 {
			q.o = append(q.o, q.i[len(q.i)-1])
			q.i = q.i[:len(q.i)-1]
		}
	}
	v := q.o[len(q.o)-1]
	q.o = q.o[:len(q.o)-1]
	return v
}
func main() {
	queue := NewQueue()
	for i := 0; i < 10; i++ {
		queue.Add(i)
	}
	for i := 0; i < 10; i++ {
		ele := queue.Del()
		fmt.Printf("%d ", ele)
	}

}
