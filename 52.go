package main

import "fmt"

func main() {
	str := "公号：代码情缘"
	// 注意byte占1个字节, 只能保存字符不能保存汉字,因为一个汉字占用3个字节
	arr1 := []byte(str) // 12
	fmt.Println(len(arr1))
	for _, v := range arr1 {
		fmt.Printf("%c", v) // lnjæ��å��æ±�
	}

	// Go语言中rune类型就是专门用于保存汉字的
	arr2 := []rune(str)
	fmt.Println(len(arr2)) // 6
	for _, v := range arr2 {
		fmt.Printf("%c", v) // lnj李南江
	}
}
