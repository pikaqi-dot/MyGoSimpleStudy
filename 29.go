package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	// 创建时利用取地址符号获取结构体变量地址
	var p1 = &Student{"lnj", 33}
	fmt.Println(p1) // &{lnj 33}

	// 通过new内置函数传入数据类型创建
	// 内部会创建一个空的结构体变量, 然后返回这个结构体变量的地址
	var p2 = new(Student)
	fmt.Println(p2) // &{ 0}
}
