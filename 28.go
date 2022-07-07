package main

import "fmt"

func main() {
	// 1.定义一个切片
	var sce []int = []int{1, 3, 5}
	// 2.打印切片的地址
	// 切片变量中保存的地址, 也就是指向的那个数组的地址 sce = 0xc0420620a0
	fmt.Printf("sce = %p\n", sce)
	fmt.Println(sce) // [1 3 5]
	// 切片变量自己的地址, &sce = 0xc04205e3e0
	fmt.Printf("&sce = %p\n", &sce)
	fmt.Println(&sce) // &[1 3 5]
	// 3.定义一个指向切片的指针
	var p *[]int
	// 因为必须类型一致才能赋值, 所以将切片变量自己的地址给了指针
	p = &sce
	// 4.打印指针保存的地址
	// 直接打印p打印出来的是保存的切片变量的地址 p = 0xc04205e3e0
	fmt.Printf("p = %p\n", p)
	fmt.Println(p) // &[1 3 5]
	// 打印*p打印出来的是切片变量保存的地址, 也就是数组的地址 *p = 0xc0420620a0
	fmt.Printf("*p = %p\n", *p)
	fmt.Println(*p) // [1 3 5]

	// 5.修改切片的值
	// 通过*p找到切片变量指向的存储空间(数组), 然后修改数组中保存的数据
	(*p)[1] = 666
	fmt.Println(sce[1])
}
