package main

import "fmt"

// 1.定义接口
type Animal interface {
	Eat()
}
type Dog struct {
	name string
	age  int
}

// 2.实现接口方法
func (d Dog) Eat() {
	fmt.Println(d.name, "正在吃东西")
}

type Cat struct {
	name string
	age  int
}

// 2.实现接口方法
func (c Cat) Eat() {
	fmt.Println(c.name, "正在吃东西")
}

// 3.对象特有方法
func (c Cat) Special() {
	fmt.Println(c.name, "特有方法")
}

func main() {
	// 1.利用接口类型保存实现了所有接口方法的对象
	var a Animal
	a = Dog{"旺财", 18}
	// 2.利用接口类型调用对象中实现的方法
	a.Eat()
	a = Cat{"喵喵", 18}
	a.Eat()
	// 3.利用接口类型调用对象特有的方法
	//a.Special() // 接口类型只能调用接口中声明的方法, 不能调用对象特有方法
	if cat, ok := a.(Cat); ok {
		cat.Special() // 只有对象本身才能调用对象的特有方法
	}
}
