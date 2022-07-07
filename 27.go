//https://github.com/pikaqi-dot/GoGuide
package main

import "fmt"

func main() {

	var arr [3]int = [3]int{1, 3, 5}
	var p *[3]int
	p = &arr
	fmt.Printf("%p\n", &arr) // 0xc0420620a0
	fmt.Printf("%p\n", p)    // 0xc0420620a0
	fmt.Println(&arr)        // &[1 3 5]
	fmt.Println(p)           // &[1 3 5]
	// 指针指向数组之后操作数组的几种方式
	// 1.直接通过数组名操作
	arr[1] = 6
	fmt.Println(arr[1])
	// 2.通过指针间接操作
	(*p)[1] = 7
	fmt.Println((*p)[1])
	fmt.Println(arr[1])
	// 3.通过指针间接操作
	p[1] = 8
	fmt.Println(p[1])
	fmt.Println(arr[1])

	// 注意点: Go语言中的指针, 不支持+1 -1和++ --操作
	//*(p + 1) = 9 // 报错
	//fmt.Println(*p++) // 报错
	//fmt.Println(arr[1])
}
