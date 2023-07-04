package genericParadigm

import (
	"fmt"
	"testing"
)

// TestGp
//
//	@Description: 测试泛型
//	@param t
func TestSliceGp(t *testing.T) {
	type Slice[T int | float64] []T
	//T是类型形参  int|float64 是类型约束  []包含的是类型形参列表  新类型 Slice[T]
	var a = make(Slice[int], 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)
}
func TestMapGp(t *testing.T) {
	type Map[KEY int | string, VALUE int | string] map[KEY]VALUE
	//KEY,VALUE 类型形参  int | string 类型约束  map[KEY]VALUE 泛型类型
	var a Map[string, int] = make(map[string]int)
	a["张三"] = 10
	fmt.Println(a)
}
func TestStructGp(t *testing.T) {
	type MyStruct[T int | string | float64] struct {
		Name string
		Data T
	}
	s1 := MyStruct[int]{
		Name: "张三",
		Data: 10,
	}
	fmt.Println(s1)
}

type p[T any] interface {
	Print()
}

type Car struct {
	brand string
}

func NewCar(brand string) *Car {
	return &Car{brand: brand}
}

func moveChe() p[string] {
	return Car{}
}

func (c Car) Print() {
	fmt.Println("!")
}

func TestInterfaceGp(t *testing.T) {
	che := moveChe()
	che.Print()
	c := NewCar("张三")
	c.Print()
	fmt.Println()
}

func TestChannelGp(t *testing.T) {
	type MyChan[T any] chan T

}

type myTypeSlice[T int | float64] []T

func (s myTypeSlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func TestSum(t *testing.T) {
	m := make(myTypeSlice[int], 0)
	for i := 0; i < 10; i++ {
		m = append(m, i)
	}
	sum := m.Sum()
	fmt.Println(sum)
}

func Add[T int | float64](a, b T) T {
	return a + b
}
func TestAdd(t *testing.T) {
	fmt.Println(Add(1, 2))
}

func MyFunc[T int | float64](a, b T) {
	fn2 := func(i, j T) T {
		return i*2 + j*2
	}
	fmt.Println(fn2(a, b))
}

func TestFunc(t *testing.T) {
	MyFunc(1, 2)
}
