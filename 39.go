package main

import "fmt"

type studier interface {
	read()
}
type Person struct {
	name string
	age  int
}

func (p Person) read() {
	fmt.Println(p.name, "正在学习")
}
func main() {
	var s studier
	s = Person{"lnj", 33}
	//s.name = "zs" // 报错, 由于s是接口类型, 所以不能访问属性
	// 2.定义一个结构体类型变量
	//var p Person
	// 不能用强制类型转换方式将接口类型转换为原始类型
	//p = Person(s) // 报错

	// 2.利用ok-idiom模式将接口类型还原为原始类型
	// s.(Person)这种格式我们称之为: 类型断言
	if p, ok := s.(Person); ok {
		p.name = "zs"
		fmt.Println(p)
	}

	// 2.通过 type switch将接口类型还原为原始类型
	// 注意: type switch不支持fallthrought
	switch p := s.(type) {
	case Person:
		p.name = "zs"
		fmt.Println(p) // {zs 33}
	default:
		fmt.Println("不是Person类型")
	}
}
