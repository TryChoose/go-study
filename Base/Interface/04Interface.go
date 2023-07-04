package main

import "fmt"

//支付方式

type payer interface {
	pay()
}
type WeChat struct {
	many float64
}
type AliPay struct {
	many float64
}

func (w *WeChat) pay() {
	fmt.Printf("使用微信方式,支付了:%.2f元\n", w.many)
}
func (a *AliPay) pay() {
	fmt.Printf("使用支付宝方式,支付了:%.2f元\n", a.many)

}
func payMethod(p payer) {
	p.pay()
}
func main() {
	fmt.Print(`1.支付宝
2.微信` + "\n")
	fmt.Printf("请选择支付方式：")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println("请输入您支付的金额:")
		var many float64
		fmt.Scan(&many)
		payMethod(&AliPay{many: many})
	} else if choice == 2 {
		fmt.Println("请输入您支付的金额:")
		var many float64
		fmt.Scan(&many)
		payMethod(&WeChat{many: many})
	}

}
