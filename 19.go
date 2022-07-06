//sync 包还提供了一个条件变量类型 sync.Cond，它可以和互斥锁或读写锁（以下统称互斥锁）组合使用，用来协调想要访问共享资源的线
//程。
//
//不过，与互斥锁不同，条件变量 sync.Cond 的主要作用并不是保证在同一时刻仅有一个线程访问某一个共享资源，而是在对应的共享资源状
//态发生变化时，通知其它因此而阻塞的线程。条件变量总是和互斥锁组合使用，互斥锁为共享资源的访问提供互斥支持，而条件变量可以就共享资
//源的状态变化向相关线程发出通知，重在「协调」。
//
//下面，我们来看看如何使用条件变量 sync.Cond。
//
//sync.Cond 是一个结构体：
//type Cond struct {
//  noCopy noCopy
//
//  // L is held while observing or changing the condition
//  L Locker
//
//  notify  notifyList
//  checker copyChecker
//}

//提供了三个方法：
//// 等待通知
//func (c *Cond) Wait() {
//  c.checker.check()
//  t := runtime_notifyListAdd(&c.notify)
//  c.L.Unlock()
//  runtime_notifyListWait(&c.notify, t)
//  c.L.Lock()
//}
//
//// 单发通知
//func (c *Cond) Signal() {
//  c.checker.check()
//  runtime_notifyListNotifyOne(&c.notify)
//}
//
//// 广播通知
//func (c *Cond) Broadcast() {
//  c.checker.check()
//  runtime_notifyListNotifyAll(&c.notify)
//}
//我们可以通过 sync.NewCond 返回对应的条件变量实例，初始化的时候需要传入互斥锁，该互斥锁实例会赋值给 sync.Cond 的 L 属性：
//locker := &sync.Mutex{}
//cond := sync.NewCond(locker)

//sync.Cond 主要实现一个条件变量，假设 goroutine A 执行前需要等待另外一个 goroutine B 的通知，那么处于等待状态的
//goroutine A 会保存在一个通知列表，也就是说需要某种变量状态的 goroutine A 将会等待（Wait）在那里，当某个时刻变量状态
//改变时，负责通知的 goroutine B 会通过对条件变量通知的方式（Broadcast/Signal）来通知处于等待条件变量的 goroutine A，
//这样就可以在共享内存中实现类似「消息通知」的同步机制。
//
//下面来看一个具体的示例。假设我们有一个读取器和一个写入器，读取器必须依赖写入器对缓冲区进行数据写入后，才可以从缓冲区中读取数据，
//写入器每次完成写入数据后，都需要通过某种通知机制通知处于阻塞状态的读取器，告诉它可以对数据进行访问，这种场景正好可以通过条件变量
//来实现：

package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// 数据 bucket
type DataBucket struct {
	buffer *bytes.Buffer  //缓冲区
	mutex *sync.RWMutex //互斥锁
	cond  *sync.Cond //条件变量
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer:     bytes.NewBuffer(buf),
		mutex: new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

// 读取器
func (db *DataBucket) Read(i int) {
	db.mutex.RLock()   // 打开读锁
	defer db.mutex.RUnlock()  // 结束后释放读锁
	var data []byte
	var d byte
	var err error
	for {
		//每次读取一个字节
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF { // 缓冲区数据为空时执行
				if string(data) != "" {  // data 不为空，则打印它
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.cond.Wait() // 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				data = data[:0]  // 将 data 清空
				continue
			}
		}
		data = append(data, d) // 将读取到的数据添加到 data 中
	}
}

// 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()   // 打开写锁
	defer db.mutex.Unlock()  // 结束后释放写锁
	//写入一个数据块
	n, err := db.buffer.Write(d)
	db.cond.Signal()  // 写入数据后通过 Signal 通知处于阻塞状态的读取器
	return n, err
}

func main() {
	db := NewDataBucket()
	go db.Read(1) // 开启读取器协程
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		db.Put([]byte(d))  // 写入数据到缓冲区
	}(1)  // 开启写入器协程
	time.Sleep(100 * time.Millisecond)
}


//这里我们使用了读写互斥锁，在读取器里面使用读锁，在写入器里面使用写锁，并且通过 defer 语句释放锁，然后在锁保护的情况下，通过条件
//变量协调读写线程：在读线程中，当缓冲区为空的时候，通过 db.cond.Wait() 阻塞读线程；在写线程中，当缓冲区写入数据的时候通过
//db.cond.Signal() 通知读线程继续读取数据。
//
//执行上述示例代码，结果如下：
//
//reader-1: data-1











