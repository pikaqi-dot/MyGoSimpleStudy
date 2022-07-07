package main

import "fmt"

type studier interface {
	read()
}
type Person struct {
	name string
	age  int
}

func (p Person) read() {
	fmt.Println(p.name, "正在学习")
}
func main() {
	// 1.定义一个抽象接口类型
	var i interface{}
	i = Person{"lnj", 33}
	// 不能调用read方法, 因为抽象接口中没有这个方法
	//i.read()
	// 2.利用ok-idiom模式将抽象接口转换为具体接口
	if s, ok := i.(studier); ok {
		// 可以调用read方法,因为studier中声明了这个方法,并且结构体中实现了这个方法
		s.read() // lnj 正在学习
	}
}
