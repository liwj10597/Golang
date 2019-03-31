package main

import "fmt"

/**
指针
*/

func main() {
	// 声明int型指针
	var point *int

	var num int
	num = 100
	// 将变量num的地址赋予指针point
	point = &num

	// 打印指针
	fmt.Println(point)
	// 打印指针指向的值
	fmt.Println(*point)

	// 改变指针所指向的值内容(前提是指针必须要有地址)
	*point = 101
	fmt.Println(*point)

	// 指针判空
	if point == nil {
		fmt.Println("指针为空")
	}
}
