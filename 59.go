package main

import "fmt"

func main() {
	// 1.创建一个管道
	mych := make(chan int, 3)
	// 2.往管道中存入数据
	mych <- 666
	mych <- 777
	mych <- 888
	// 3.遍历管道
	// 第一次遍历i等于0, len = 3,
	// 第二次遍历i等于1, len = 2
	// 第三次遍历i等于2, len = 1
	//for i:=0; i<len(mych); i++{
	//	fmt.Println(<-mych) // 输出结果不正确
	//}

	// 3.写入完数据之后先关闭管道
	// 注意点: 管道关闭之后只能读不能写
	close(mych)
	//mych<- 999 // 报错

	// 4.遍历管道
	// 利用for range遍历, 必须先关闭管道, 否则会报错
	//for value := range mych{
	//	fmt.Println(value)
	//}

	// close主要用途:
	// 在企业开发中我们可能不确定管道有还没有有数据, 所以我们可能一直获取
	// 但是我们可以通过ok-idiom模式判断管道是否关闭, 如果关闭会返回false给ok
	for {
		if num, ok := <-mych; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
	fmt.Println("数据读取完毕")
}
