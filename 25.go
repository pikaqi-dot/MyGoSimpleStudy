//sync.Pool 是一个临时对象池，可用来临时存储对象，下次使用时从对象池中获取，避免重复创建对象。相应的，该类型提供了 Put 和 Get 方法，
//分别对临时对象进行存储和获取。我们可以把 sync.Pool 看作存放可重复使用值的容器，由于 Put 方法支持的参数类型是空接口 interface{}，因
//此这个值可以是任何类型，对应的，Get 方法返回值类型也是 interface{}。当我们通过 Get 方法从临时对象池获取临时对象后，会将原来存放在里面的
//对象删除，最后再返回这个对象，而如果临时对象池中原来没有存储任何对象，调用 Get 方法时会通过对象池的 New 字段对应函数创建一个新值并返回（这
//个 New 字段需要在初始化临时对象池时指定，否则对象池为空时调用 Get 方法返回的可能就是 nil），从而保证无论临时对象池中是否存在值，始终都能返回
//结果。



package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello,World!"
		},
	}
	value := "Hello,学院君!"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
//在这段代码中，我们首先声明并初始化了一个临时对象池 pool，并定义了 New 字段，这是一个 func() interface{} 类型的函数，然后我们通过
//Put 方法存储一个字符串对象到 pool，最后通过 Get 方法获取该对象并打印，当我们再次在 pool 实例上调用 Get 方法时，会发现存储的字符串已
//经不存在，而是通过 New 字段对应函数返回的字符串对象。
//
//此外，我们还可以利用 sync.Pool 的特性在多协程并发执行场景下实现对象的复用，因为 sync.Pool 本身是并发安全地，我们可以在程序开始执行时全局
//唯一初始化 Pool 对象，然后在并发执行的协程之间通过这个临时对象池来存储和获取对象：


