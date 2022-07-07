package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 1.创建管道
	myCh := make(chan int, 5)
	exitCh := make(chan bool)

	// 2.生成数据
	go func() {
		for i := 0; i < 10; i++ {
			myCh <- i
			time.Sleep(time.Second * 3)
		}
	}()

	// 3.获取数据
	go func() {
		for {
			select {
			case num := <-myCh:
				fmt.Println(num)
			case <-time.After(time.Second * 2):
				exitCh <- true
				runtime.Goexit()
			}
		}
	}()

	<-exitCh
	fmt.Println("程序结束")
}
