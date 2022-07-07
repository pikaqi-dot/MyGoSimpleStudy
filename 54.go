package main

import (
	"fmt"
	"runtime"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在唱歌")
		// Gosched使当前go程放弃处理器，以让其它go程运行。
		// 它不会挂起当前go程，因此当前go程未来会恢复执行
		runtime.Gosched()
	}
}
func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("我在跳舞---")
		runtime.Gosched()
	}
}

func main() {

	go sing()
	go dance()
	for {

	}
}
