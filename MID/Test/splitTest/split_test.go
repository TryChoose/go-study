package splitTest

import "testing"

//func TestSplit(t *testing.T) {
//	got:= Split("abcdeabc", "de")
//	want:=[]string{"abc","abc"}
//	if !reflect.DeepEqual(got,want){
//		t.Errorf("want:%v but got:%v\n",want,got)
//	}
//}

type testCase struct {
	got  string
	sep  string
	want []string
}

//测试组
//func TestSplit(t *testing.T){
//	test:=[]testCase{
//		{got: "abc",sep: "b",want: []string{"a","c"}},
//		{got: "张三了李四",sep: "了",want: []string{"张三","四"}},
//	}
//	//遍历切片
//	for _,tc:= range test {
//		res := Split(tc.got, tc.sep)
//		if !reflect.DeepEqual(res,tc.want){
//			t.Fatalf("测试用例不通过 expected:%#v,res:%#v",tc.want,res)
//		}
//	}
//}

//子测试
//func TestSplit(t *testing.T) {
//	test := map[string]testCase{
//		"case_1": {got: "abc", sep: "b", want: []string{"a", "c"}},
//		"case_2": {got: "张三了李四", sep: "了", want: []string{"张三", "李四"}},
//		"case_3": {got: "ABCDEFGABCDEF", sep: "EF", want: []string{"ABCD", "GABCD",""}},
//	}
//	//遍历切片
//	for name, tc := range test {
//		t.Run(name, func(t *testing.T) {
//			res := Split(tc.got, tc.sep)
//			if !reflect.DeepEqual(res, tc.want) {
//				t.Fatalf("测试用例不通过 expected:%#v,res:%#v", tc.want, res)
//			}
//		})
//	}
//}

// 基准测试
//func BenchmarkSplit(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Split("a:b:c", ":")
//	}
//}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
func BenchmarkFib50(b *testing.B) { benchmarkFib(b, 50) }
