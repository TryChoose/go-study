package main

import (
	"errors"
	"fmt"
)

// Queue 队列
type Queue struct {
	//队列的长度
	Length int
	//队列的容量
	Capacity int
	//队列的数据
	Data []interface{}
}

// NewQueue 创建一个队列
func NewQueue(length int, capacity int, data []interface{}) *Queue {
	return &Queue{Length: length, Capacity: capacity, Data: data}
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.Length == 0
}

// IsFull 判断队列是否已满
func (q *Queue) IsFull() bool {
	if q.Length == q.Capacity {
		q.AutoExpansion()
	}
	return false
}

// Enqueue 入队
func (q *Queue) Enqueue(data interface{}) error {
	if q.IsFull() {
		return errors.New("队列已满")
	}
	q.Data = append(q.Data, data)
	q.Length++
	return nil
}

// Dequeue 出队
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("队列为空")
	}
	data := q.Data[0]
	q.Data = q.Data[1:]
	q.Length--
	return data, nil
}

// AutoExpansion 自动触发队列扩容
func (q *Queue) AutoExpansion() {
	if q.Length == q.Capacity {
		q.Capacity *= 2
	}
}

// 释放队列
func (q *Queue) Close() {
	q.Data = []interface{}{}
	q.Length = 0
	q.Capacity = 0
}
func main() {
	q := NewQueue(10, 20, []interface{}{})
	for i := 0; i < 10; i++ {
		err := q.Enqueue(i)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	dequeue, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dequeue)
	}
	fmt.Println(q.Data)
	defer q.Close()
}
