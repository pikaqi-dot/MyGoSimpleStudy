//sync.RWMutex
//Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个
//goroutine 释放该 Mutex，不管是读操作还是写操作都会阻塞，但其实我们知道为了提升性能，读操作往往是不需要阻塞的，因此 sync
//包提供了 RWMutex 类型，即读/写互斥锁，简称读写锁，这是一个是单写多读模型。
//
//sync.RWMutex 分读锁和写锁，会对读操作和写操作区分对待，在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine
//可同时获取读锁，读锁调用 RLock() 方法开启，通过 RUnlock 方法释放；而写锁会阻止任何其他 goroutine（无论读和写）进来，整
//个锁相当于由该 goroutine 独占，和 Mutex 一样，写锁通过 Lock 方法启用，通过 Unlock 方法释放，从 RWMutex 的底层实现看
//实际上是组合了 Mutex：
//type RWMutex struct {
//    w Mutex
//    writerSem uint32
//    readerSem uint32
//    readerCount int32
//    readerWait int32
//}

//同样，使用 RWMutex 时，任何一个 Lock() 或 RLock() 均需要保证有对应的 Unlock() 或 RUnlock() 调用与之对应，否则可能导
//致等待该锁的所有 goroutine 处于阻塞状态，甚至可能导致死锁，比如我们可以通过 RWMutex 重构上面示例代码的锁，效果完全一样：

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

