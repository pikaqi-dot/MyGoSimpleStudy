package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在唱歌")
		time.Sleep(time.Millisecond)
	}
}
func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在跳舞---")
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// 串行: 必须先唱完歌才能跳舞
	//sing()
	//dance()

	// 并行: 可以边唱歌, 边跳舞
	// 注意点: 主线程不能死, 否则程序就退出了
	go sing()  // 开启一个协程
	go dance() // 开启一个协程
	for {

	}
}
