package main

import "fmt"

type Object struct {
	life int
}
type Person struct {
	Object
	name string
	age  int
}
type Student struct {
	Person
	score int
}

func main() {
	s := Student{Person{Object{77}, "zs", 33}, 99}
	//fmt.Println(s.Person.Object.life)
	//fmt.Println(s.Person.life)
	fmt.Println(s.life) // 三种方式都可以
	//fmt.Println(s.Person.name)
	fmt.Println(s.name) // 两种方式都能访问
	//fmt.Println(s.Person.age)
	fmt.Println(s.age) // 两种方式都能访问
	fmt.Println(s.score)
}
