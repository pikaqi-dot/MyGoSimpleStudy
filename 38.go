package main

import "fmt"

// 1.定义一个接口
type usber interface {
	start()
	stop()
}

// 2.自定义int类型
type integer int

// 2.实现接口中的所有方法
func (i integer) start() {
	fmt.Println("int类型实现接口")
}
func (i integer) stop() {
	fmt.Println("int类型实现接口")
}
func main() {
	var i integer = 666
	i.start() // int类型实现接口
	i.stop()  // int类型实现接口
}
