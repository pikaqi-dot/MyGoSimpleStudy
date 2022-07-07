package main

import "fmt"

var myCh1 = make(chan int, 5)
var myCh2 = make(chan int, 0)

func main() {
	// 有缓冲管道
	// 只写入, 不读取不会报错
	myCh1 <- 1
	myCh1 <- 2
	myCh1 <- 3
	myCh1 <- 4
	myCh1 <- 5
	fmt.Println("len =", len(myCh1), "cap =", cap(myCh1))

	// 无缓冲管道
	// 只有两端同时准备好才不会报错
	go func() {
		fmt.Println(<-myCh2)
	}()
	fmt.Println("len =", len(myCh2), "cap =", cap(myCh2))
	// 只写入, 不读取会报错
	myCh2 <- 1
	fmt.Println("len =", len(myCh2), "cap =", cap(myCh2))
	// 写入之后在同一个线程读取也会报错
	//fmt.Println(<-myCh2)
	// 在主程中先写入, 在子程中后读取也会报错
	//go func() {
	//	fmt.Println(<-myCh2)
	//}()
}
