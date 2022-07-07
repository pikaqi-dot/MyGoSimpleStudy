package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("123")
		// 退出当前协程
		//runtime.Goexit()
		// 退出当前函数
		//return
		test()
		fmt.Println("456")
	}()
	for {

	}
}

func test() {
	fmt.Println("abc")
	// 只会结束当前函数, 协程中的其它代码会继续执行
	//return
	// 会结束整个协程, Goexit之后整个协程中的其它代码不会执行
	runtime.Goexit()
	fmt.Println("def")
}
