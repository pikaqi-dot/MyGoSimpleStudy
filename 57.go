//上述实现并发的代码中为了保持主线程不挂掉, 我们都会在最后写上一个死循环或者写上一个定时器来实现等待goroutine执行完毕
//上述实现并发的代码中为了解决生产者消费者资源同步问题, 我们利用加锁来解决, 但是这仅仅是一对一的情况, 如果是一对多或者多对多, 上述代码
//还是会出现问题
//综上所述, 企业开发中需要一种更牛X的技术来解决上述问题, 那就是管道(Channel)
package main

import "fmt"

func main() {
	// 1.声明一个管道
	var mych chan int
	// 2.初始化一个管道
	mych = make(chan int, 3)
	// 3.查看管道的长度和容量
	fmt.Println("长度是", len(mych), "容量是", cap(mych))
	// 4.像管道中写入数据
	mych <- 666
	fmt.Println("长度是", len(mych), "容量是", cap(mych))
	// 5.取出管道中写入的数据
	num := <-mych
	fmt.Println("num = ", num)
	fmt.Println("长度是", len(mych), "容量是", cap(mych))
}
