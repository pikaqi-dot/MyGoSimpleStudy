package main

import "fmt"

type A interface {
	fna()
}
type B interface {
	fnb()
}
type C interface {
	A // 嵌入A接口
	B // 嵌入B接口
	fnc()
}
type Person struct{}

func (p Person) fna() {
	fmt.Println("实现A接口中的方法")
}
func (p Person) fnb() {
	fmt.Println("实现B接口中的方法")
}
func (p Person) fnc() {
	fmt.Println("实现C接口中的方法")
}
func main() {
	p := Person{}
	p.fna() // 实现A接口中的方法
	p.fnb() // 实现B接口中的方法
	p.fnc() // 实现C接口中的方法
}
