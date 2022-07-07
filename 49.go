package main

import "fmt"

func div(a, b int) (res int) {
	if b == 0 {
		//err = errors.New("除数不能为0")
		panic("除数不能为0")
	} else {
		res = a / b
	}
	return
}
func main() {
	// panic异常会沿着调用堆栈向外传递, 所以也可以在外层捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 除数不能为0
		}
	}()
	div(10, 0)
}
