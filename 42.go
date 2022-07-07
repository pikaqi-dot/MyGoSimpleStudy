package main

import "fmt"

type Person struct {
	name string // 属性重名
	age  int
}
type Student struct {
	Person
	name  string // 属性重名
	score int
}

func main() {
	s := Student{Person{"zs", 18}, "ls", 99}

	fmt.Println(s.Person.name) // zs
	fmt.Println(s.name)        // ls
	//fmt.Println(s.Person.age)
	fmt.Println(s.age) // 两种方式都能访问
	fmt.Println(s.score)
}
