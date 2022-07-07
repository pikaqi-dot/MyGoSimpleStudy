package main

import "fmt"

// 1.定义一个接口
type usber interface {
	start()
	stop()
}
type Computer struct {
	name  string
	model string
}

// 2.实现接口中的所有方法
func (cm Computer) start() {
	fmt.Println("启动电脑")
}
func (cm Computer) stop() {
	fmt.Println("关闭电脑")
}

type Phone struct {
	name  string
	model string
}

// 2.实现接口中的所有方法
func (p Phone) start() {
	fmt.Println("启动手机")
}
func (p Phone) stop() {
	fmt.Println("关闭手机")
}

// 3.使用接口定义的方法
func working(u usber) {
	u.start()
	u.stop()
}
func main() {
	cm := Computer{"戴尔", "F1234"}
	working(cm) // 启动电脑 关闭电脑

	p := Phone{"华为", "M10"}
	working(p) // 启动手机 关闭手机
}
