package main

import "fmt"

func test1(x, y int) {
	// 多个异常,只有第一个会被捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 异常A
		}
	}()
	panic("异常A") // 相当于return, 后面代码不会继续执行
	panic("异常B")
}
func main() {
	test1(10, 0)
}
