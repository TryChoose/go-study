package stack

import "fmt"

// Stack
// @Description: 实现栈
type Data interface{}
type Stack struct {
	DataMENT []Data
}

// NewStack
//
//	@Description: 构造一个栈
//	@param data
//	@return *Stack

func (st *Stack) NewStack(dataMENT []Data) *Stack {
	return &Stack{DataMENT: dataMENT}
}

// Length
//
//	@Description: 返回该栈的长度
//	@receiver s
//	@param stack
//	@return int
func (st *Stack) Length(stack *Stack) int {
	return len(stack.DataMENT)
}

// isEmpty
//
//	@Description: 判断是否为空
//	@receiver s
//	@param stack
//	@return bool
func (st *Stack) IsEmpty(stack *Stack) bool {
	if len(stack.DataMENT) == 0 {
		return true
	}
	return false
}
func (st *Stack) Push(stack *Stack, data Data) {
	stack.DataMENT = append(stack.DataMENT, data)
}

func (st *Stack) Pop(stack *Stack) *Data {
	if !st.IsEmpty(stack) {
		fmt.Println("栈的空间没有元素了!")
		return nil
	}
	s := stack.DataMENT[len(stack.DataMENT)-1]
	st.DataMENT = st.DataMENT[0 : len(stack.DataMENT)-1]
	return &s
}
