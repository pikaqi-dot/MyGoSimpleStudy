package main

import "fmt"

func main() {
	// 1.定义一个双向管道
	var myCh chan int = make(chan int, 5)

	// 2.将双向管道转换单向管道
	var myCh2 chan<- int
	myCh2 = myCh
	fmt.Println(myCh2)
	var myCh3 <-chan int
	myCh3 = myCh
	fmt.Println(myCh3)
	myCh <- 1
	myCh <- 2
	myCh <- 3

	// 3.只写管道,只能写, 不能读
	myCh2 <- 666
	close(myCh2)
	//	fmt.Println(<-myCh2)
	// 4.只读管道, 只能读,不能写
	//myCh3<-666
	for {
		if num, ok := <-myCh3; ok {
			fmt.Println("xxxxxxx")
			fmt.Println(num)
		} else {
			fmt.Println("没内容")
			break
			//return
		}
	}
	//close(myCh)
	// 注意点: 管道之间赋值是地址传递, 以上三个管道底层指向相同容器
	// 3.双向管道,可读可写

	//close(myCh)
	fmt.Println("--------------")
	for {

		if num, ok := <-myCh; ok {
			fmt.Println(num)
		} else {
			fmt.Println("没内容")
			break
		}
	}

}
