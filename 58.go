package main

import "fmt"

func main() {
	// 1.声明一个管道
	var mych chan int
	// 2.初始化一个管道
	mych = make(chan int, 3)

	// 注意点: 管道中只能存放声明的数据类型, 不能存放其它数据类型
	mych <- 3.14

	// 注意点: 管道中如果已经没有数据,
	// 并且检测不到有其它协程再往管道中写入数据, 那么再取就会报错
	num = <-mych
	fmt.Println("num = ", num)

	// 注意点: 如果管道中数据已满, 再写入就会报错
	mych <- 666
	mych <- 777
	mych <- 888
	mych <- 999
}
