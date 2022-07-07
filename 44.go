package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 父类方法
func (p Person) say() {
	fmt.Println("name is ", p.name, "age is ", p.age)
}

type Student struct {
	Person
	score float32
}

func main() {
	stu := Student{Person{"zs", 18}, 59.9}
	stu.say()
}
