package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	var p = &Student{}
	// 方式一: 传统方式操作
	// 修改结构体中某个属性对应的值
	// 注意: 由于.运算符优先级比*高, 所以一定要加上()
	(*p).name = "lnj"
	// 获取结构体中某个属性对应的值
	fmt.Println((*p).name) // lnj

	// 方式二: 通过Go语法糖操作
	// Go语言作者为了程序员使用起来更加方便, 在操作指向结构体的指针时可以像操作接头体变量一样通过.来操作
	// 编译时底层会自动转发为(*p).age方式
	p.age = 33
	fmt.Println(p.age) // 33
}
