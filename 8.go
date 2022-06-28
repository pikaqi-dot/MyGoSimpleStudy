//Go 语言还支持通过 select 分支语句选择指定分支代码执行，select 语句和之前介绍的 switch 语句语法结构类似，不同之处在于
//select 的每个 case 语句必须是一个通道操作，要么是发送数据到通道，要么是从通道接收数据，此外 select 语句也支持 default
//分支：
//
//
//select {
//
//    case <-chan1:
//
//        // 如果从 chan1 通道成功接收数据，则执行该分支代码
//
//    case chan2 <- 1:
//
//        // 如果成功向 chan2 通道成功发送数据，则执行该分支代码
//
//    default:
//7
//        // 如果上面都没有成功，则进入 default 分支处理流程
//8
//}
//注：Go 语言的 select 语句借鉴自 Unix 的 select() 函数，在 Unix 中，可以通过调用 select() 函数来监控一系列的文件句柄，
//一旦其中一个文件句柄发生了 IO 动作，该 select() 调用就会被返回（C 语言中就是这么做的），后来该机制也被用于实现高并发的 Socket
//服务器程序。Go 语言直接在语言级别支持 select 关键字，用于处理并发编程中通道之间异步 IO 通信问题。
//
//可以看出，select 不像 switch，case 后面并不带判断条件，而是直接去查看 case 语句，每个 case 语句都必须是一个面向通道的操作，
//比如上面的示例代码中，第一个 case 试图从 chan1 接收数据并直接忽略读到的数据，第二个 case 试图向 chan2 通道发送一个整型数据
//需要注意的是这两个 case 的执行不是 if...else... 那种先后关系，而是会并发执行，然后 select 会选择先操作成功返回的那个 case
//分支去执行，如果两者同时返回，则随机选择一个执行，如果这两者都没有返回，则进入 default 分支，这里也不会出现阻塞，如果 chan1 通道为空，
//或者 chan2 通道已满，就会立即进入 default 分支，但是如果没有 default 语句，则会阻塞直到某个通道操作成功。
//
//因此，借助 select 语句我们可以在一个协程中同时等待多个通道达到就绪状态
//这些通道操作是并发的，任何一个操作成功，就会进入该分支执行代码，否则程序就会处于挂起状态，如果要实现非阻塞操作，可以引入 default
//语句。
//
//下面我们基于 select 语句来实现一个简单的示例代码：
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3) // 随机生成0-2之间的数字
	fmt.Printf("随机索引/数值: %d\n", index)
	chs[index] <- index // 向通道发送随机数字

	// 哪一个通道中有值，哪个对应的分支就会被执行
	select {
	case <-chs[0]:
		fmt.Println("第一个条件分支被选中")
	case <-chs[1]:
		fmt.Println("第二个条件分支被选中")
	case num := <-chs[2]:
		fmt.Println("第三个条件分支被选中:", num)
	default:
		fmt.Println("没有分支被选中")
	}
}

//如果我们将 chs[index] <- index 这一行注释掉，则打印结果如下：
//
//
//随机索引/数值: 2
//
//没有分支被选中
//另外，需要注意的是，select 语句只能对其中的每一个 case 表达式各求值一次，如果我们想连续操作其中的通道的话，需要
//通过在 for 语句中嵌入 select 语句的方式来实现
