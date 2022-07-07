package main

import "fmt"

type aer interface {
	fna()
}
type ber interface {
	aer
	fnb()
}

// Person实现了超集接口所有方法
type Person struct{}

func (p Person) fna() {
	fmt.Println("实现A接口中的方法")
}
func (p Person) fnb() {
	fmt.Println("实现B接口中的方法")
}

// Student实现了子集接口所有方法
type Student struct{}

func (p Student) fna() {
	fmt.Println("实现A接口中的方法")
}
func main() {
	var i ber
	// 子集接口变量不能转换为超集接口变量
	//i = Student{}
	fmt.Println(i)
	var j aer
	// 超集接口变量可以自动转换成子集接口变量,
	j = Person{}
	fmt.Println(j)
}
