package main

import "fmt"

func main() {
	//var n float64
	//var m int
	//var s float64
	//fmt.Scanf("%f%d", &n, &m)
	//for m != 0 {
	//	for s = 0; m > 0; m-- {
	//		s += n
	//		n = math.Sqrt(n)
	//	}
	//	fmt.Printf("%.2f\n", s)
	//	fmt.Scanf("%f%d", &n, &m)
	//}
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	for {
		res := f1(n, m)
		if len(res) == 0 {
			fmt.Println("no")
		} else {
			for _, v := range res {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
		_, err := fmt.Scanf("%d%d", &n, &m)
		if err != nil {
			break
		}
	}
}

func f1(n, m int) (res []int) {
	for i := n; i <= m; i++ {
		a := i % 10
		b := (i / 10) % 10
		c := i / 100 % 10
		if a*a*a+b*b*b+c*c*c == i {
			res = append(res, i)
		}
	}
	return
}
