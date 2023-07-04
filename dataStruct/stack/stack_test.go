package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack
	fmt.Println(s.IsEmpty(&s))
	s.Push(&s, 1)
	fmt.Println(s)
	fmt.Println(s.Pop(&s))
}
