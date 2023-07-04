package main

import "fmt"

type Stack struct {
	A []int
}

func NewStack(a []int) *Stack {
	return &Stack{A: a}
}

// 出栈
func (Stack *Stack) Pop1() int {
	x := Stack.A[len(Stack.A)-1]
	//Stack.A = Stack.A[:len(Stack.A)-1]
	return x
}
func (Stack *Stack) Pop2() int {
	x := Stack.A[len(Stack.A)-1]
	Stack.A = Stack.A[:len(Stack.A)-1]
	return x
}

// 压栈
func (Stack *Stack) Push(x int) {
	Stack.A = append(Stack.A, x)
}

// 返回栈中最小元素
func (Stack *Stack) getMin() int {
	return Stack.A[len(Stack.A)-1]
}

// 比较大小
func cmp(a, b int) bool {
	if a > b {
		return false
	}
	return true
}
func queue() {
	q1 := NewStack([]int{})
	q2 := NewStack([]int{})
	var y int
	for i := 0; i < 5; i++ {
		fmt.Scan(&y)
		q1.Push(y)
	}
	for i := 0; i < 5; i++ {
		q2.Push(q1.Pop2())
	}
	fmt.Println(q1, q2)
}
func stack() {
	s1 := NewStack([]int{})
	s2 := NewStack([]int{})
	var x int
	for i := 0; i < 6; i++ {
		fmt.Scan(&x)
		s1.Push(x)
		if len(s2.A) == 0 {
			s2.Push(x)
		} else {
			if cmp(s1.Pop1(), s2.Pop1()) {
				s2.Push(x)
			}
		}
	}
	fmt.Println(s2.getMin())
}

// 使用递归函数逆序栈
func getAndRemove(s *Stack) int {
	r := s.Pop2()
	if len(s.A) == 0 {
		return r
	} else {
		l := getAndRemove(s)
		s.Push(r)
		return l
	}
}
func reverse(s *Stack) {
	if len(s.A) == 0 {
		return
	}
	i := getAndRemove(s)
	reverse(s)
	s.Push(i)
}
func main() {
	//queue()
	//stack()
	s1 := NewStack([]int{1, 2, 3, 4, 5, 6})
	reverse(s1)
	fmt.Println(s1.A)

}
