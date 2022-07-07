package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 接收者是一个变量
func (p Person) setName(name string) {
	p.name = name
}

// 接收者是一个指针
func (p *Person) setAge(age int) {
	p.age = age
}
func main() {
	per := Person{"lnj", 33}
	fmt.Println(per) // {lnj 33}
	// 值传递, 方法内部修改不会影响方法外部
	per.setName("zs")
	fmt.Println(per) // {lnj 33}
	p := &per
	// 地址传递, 方法内部修改会影响方法外部
	(*p).setAge(18)
	fmt.Println(per) // {lnj 18}
}
