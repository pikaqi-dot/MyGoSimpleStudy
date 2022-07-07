package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	Person // 学生继承了人的特性
	score  int
}
type Teacher struct {
	Person // 老师继承了人的特性
	Title  string
}

func main() {
	s := Student{Person{"lnj", 18}, 99}
	//fmt.Println(s.Person.name)
	fmt.Println(s.name) // 两种方式都能访问
	//fmt.Println(s.Person.age)
	fmt.Println(s.age) // 两种方式都能访问
	fmt.Println(s.score)
}
