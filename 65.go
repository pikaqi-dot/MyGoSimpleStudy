package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建管道
	var myCh = make(chan int)
	var exitCh = make(chan bool)

	// 生产数据
	go func() {
		for i := 0; i < 10; i++ {
			myCh <- i
			time.Sleep(time.Second)
		}
		//close(myCh)
		exitCh <- true
	}()

	// 读取数据
	for {
		fmt.Println("读取代码被执行了")
		select {
		case num := <-myCh:
			fmt.Println("读到了", num)
		case <-exitCh:
			fmt.Println("结束了")
			//break // 没用, 跳出的是select。要用return结束函数执行！
			return
		}
		fmt.Println("-----------")
	}
}
