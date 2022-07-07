package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 定义一个方法
func (p Person) say() {
	fmt.Println("my name is", p.name, "my age is", p.age)
}

// 定义一个函数
func test(p Person) {
	fmt.Println("my name is", p.name, "my age is", p.age)
}
func main() {
	per := Person{"lnj", 33}
	per.say() // my name is lnj my age is 33
	test(per) // my name is lnj my age is 33
}
