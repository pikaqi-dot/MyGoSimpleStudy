package main

import "fmt"

func test2(x, y int) {
	// 如果有异常写在defer中, 并且其它异常写在defer后面, 那么只有defer中的异常会被捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 异常A
		}
	}()

	defer func() {
		panic("异常B")
	}()
	panic("异常A")
}
func main() {
	test2(10, 0)
}
