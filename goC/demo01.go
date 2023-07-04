/**
 * @Time: 2023/6/23 22:45
 * @Author: yncn90@gmail.com
 * @File: demo01.go demo01
 * @Software: GoLand main.go
 */

package main

/*#include<stdio.h>
int add(int a, int b) {
   return a + b;
}
//冒泡排序
int sort(int a[], int n) {
	   int i, j, temp;
	   for (i = 0; i < n - 1; i++) {
		   for (j = 0; j < n - 1 - i; j++) {
			   if (a[j] > a[j + 1]) {
				   temp = a[j];
				   a[j] = a[j + 1];
				   a[j + 1] = temp;
			   }
		   }
	   }
	   return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	c := C.add(1, 2)
	a := []int{0, 7, 4, 5, 1}
	pointer := unsafe.Pointer(&a[0])
	C.sort((*C.int)(pointer), 5)
	fmt.Println(c)
}
