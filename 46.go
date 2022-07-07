package main

import "fmt"

func div(a, b int) (res int) {
	if b == 0 {
		//一旦传入的除数为0, 程序就会终止
		panic("除数不能为0")
	} else {
		res = a / b
	}
	return
}
func main() {
	res := div(10, 0)
	fmt.Println(res)
}
