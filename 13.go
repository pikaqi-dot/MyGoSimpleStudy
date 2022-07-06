//我们前面反复强调，在 Go 语言并发编程中，倡导「使用通信共享内存，不要使用共享内存通信」，而这个通信的媒介就是我们前面花大量篇幅介绍的通
//道（Channel），通道是线程安全的，不需要考虑数据冲突问题，面对并发问题，我们始终应该优先考虑使用通道，它是 first class 级别
//的，但是纵使有主角光环加持，通道也不是万能的，它也需要配角，这也是共享内存存在的价值，其他语言中主流的并发编程都是通过共享内存实现
//的，共享内存必然涉及并发过程中的共享数据冲突问题，而为了解决数据冲突问题，Go 语言沿袭了传统的并发编程解决方案 —— 锁机制，这些锁都
//位于 sync 包中。
//
//我们在通过共享内存实现并发通信这篇教程中已经用到了 sync 包提供的 Mutex 锁，锁的作用都是为了解决并发情况下共享数据的原子操作和最
//终一致性问题，在系统介绍 sync 包提供的各种锁之前，我们先来聊聊什么情况下需要用到锁。
//竞态条件与同步机制
//一旦数据被多个线程共享，那么就很可能会产生争用和冲突的情况，这种情况也被称为竞态条件（race condition），这往往会破坏共享数据的一致
//性。举个例子，
//同时有多个线程连续向同一个缓冲区写入数据块，如果没有一个机制去协调这些线程的写入操作的话，那么被写入的数据块就很可能会出现错乱。比如，
//学院君的支付宝账户余额还有 500 元，代表银行自动转账的线程 A 正在向账户转入 3000 元本月工资，同时代表花呗自动扣费的线程 B 正在从
//账户余额扣除 2000 元还上个月的花呗账单。假设用 money 标识账户余额，那么初始值 money = 500，线程 A 的操作就等价于
//money = money + 3000，线程 B 的操作就等价于 money = money - 2000，我们本来期望的结果是 money = 1500，但是现在线程 A
//和线程 B 同时对 money 进行读取和写入，所以他们拿到的 money 都是 500，如果线程 A 后执行完毕，那么 money = 3500，如果线程 B
//后执行完毕，那么 money = 0（扣除所有余额，花呗欠款1500），这就出现了和预期结果不一致的现象，我们说，这个操作破坏了数据的一致性。
//
//在这种情况下，我们就需要采取一些措施来协调它们对共享数据的修改，这通常就会涉及到同步操作。一般来说，同步的用途有两个，一个是避免多个
//线程在同一时刻操作同一个数据块，另一个是协调多个线程避免它们在同一时刻执行同一个代码块。但是目的是一致的，那就是保证共享数据原子操作
//和一致性。
//
//由于这样的数据块和代码块的背后都隐含着一种或多种资源（比如存储资源、计算资源、I/O 资源、网络资源等等），所以我们可以把它们看做是共享
//资源。我们所说的同步其实就是在控制多个线程对共享资源的访问：一个线程在想要访问某一个共享资源的时候，需要先申请对该资源的访问权限，并
//且只有在申请成功之后，访问才能真正开始；而当线程对共享资源的访问结束时，它还必须归还对该资源的访问权限，若要再次访问仍需申请。
//
//你可以把这里所说的访问权限想象成一块令牌，线程一旦拿到了令牌，就可以进入指定的区域，从而访问到资源，而一旦线程要离开这个区域了，就
//需要把令牌还回去，绝不能把令牌带走。或者我们把共享资源看作是有锁的资源，当某个线程获取到共享资源的访问权限后，给资源上锁，这样，其
//他线程就不能访问它，直到该线程执行完毕，释放锁，这样其他线程才能通过竞争获取对资源的访问权限，依次类推。
//
//这样一来，我们就可以保证多个并发运行的线程对这个共享资源的访问是完全串行的，只要一个代码片段需要实现对共享资源的串行化访问，就可以被视
//为一个临界区（critical section），也就是我刚刚说的，由于要访问到资源而必须进入的那个区域。
//
//比如，在前面举的那个例子中，实现了账户余额写入操作的代码就组成了一个临界区。临界区总是需要通过同步机制进行保护的，否则就会产生竞态
//条件，导致数据不一致。
//
//sync.Mutex
//在 Go 语言中，可供我们选择的同步工具并不少。其中，最重要且最常用的同步工具当属互斥量（mutual exclusion，简称 mutex），sync 包
//中的 Mutex 就是与其对应的类型，该类型的值可以被称为互斥锁。一个互斥锁可以被用来保护一个临界区，我们可以通过它来保证在同一时刻只有
//一个 goroutine 处于该临界区之内，回到我们通过共享内存实现并发通信这篇教程中的示例：

package main

import (
"fmt"
"runtime"
"sync"
"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

//每当有 goroutine 想进入临界区时，都需要先对它进行锁定，并且，每个 goroutine 离开临界区时，都要及时地对它进行解锁，锁定和解锁
//操作分别通过互斥锁 sync.Mutex 的 Lock 和 Unlock 方法实现。使用互斥锁的时候有以下注意事项：
//
//不要对尚未锁定或者已解锁的互斥
////不要忘记解锁互斥锁，必要时使用 defer
//不要重复锁定互斥锁；语句；锁解锁；
//不要在多个函数之间直接传递互斥锁。
