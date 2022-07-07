package main

import "fmt"

// 1.定义一个结构体
type Person struct {
	name string
	age  int
}

// 2.定义一个函数, 并将这个函数和Person结构体绑定在一起
func (p Person) say() {
	fmt.Println("my name is", p.name, "my age is", p.age)
}
func main() {
	// 3.创建一个结构体变量
	per := Person{"lnj", 33}
	// 4.利用结构体变量调用和结构体绑定的方法
	// 调用时会自动将调用者(per)传递给方法的接收者(p)
	// 所以可以在方法内部通过p方法结构体变量的属性
	per.say()
}
