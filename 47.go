package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		// 一旦传入的除数为0, 就会返回error信息
		err = errors.New("除数不能为0")
	} else {
		res = a / b
	}
	return
}
func main() {
	//res, err := div(10, 5)
	res, err := div(10, 0)
	if err != nil {
		fmt.Println(err) // 除数不能为0
	} else {
		fmt.Println(res) // 2
	}
}
