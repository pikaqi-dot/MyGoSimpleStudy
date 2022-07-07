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
	// 方式一: 先拿到指针,然后再通过指针调用
	p := &per
	(*p).setAge(18)
	fmt.Println(per) // {lnj 18}
	// 方式二: 直接利用变量调用, 底层会自动获取变量地址传递给接收者
	per.setAge(66)
	fmt.Println(per) // {lnj 66}
}
