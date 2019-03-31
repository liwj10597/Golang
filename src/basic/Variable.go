package main

import "fmt"

/**
变量
*/

// 合并var 通常用于声明全局变量
var (
	first_name string
	last_name  string
)

// 全局变量允许声明但不使用
var globalvariable int

func main() {
	// 打印
	fmt.Println("Hello GO")

	// 通过var声明变量并赋值
	var id int = 1001
	fmt.Println(id)
	// 不指定类型，通过值来反推类型
	var id2 = 2002
	fmt.Println(id2)

	// 不通过var来赋值，需要满足两点：1.name不能被声明；2.使用 := 这种形式只能用在函数内，不能用于全局变量
	name := "aubrey"
	fmt.Print(name)

	var corp, address, position string
	corp, address, position = "jollychic", "祥茂路99号", "88"
	fmt.Println(corp, address, position)

	var part, laungh, happy = 1, "li", 'a'
	fmt.Println(part, laungh, happy)

	first_name = "li"
	last_name = "jun"
	fmt.Println(first_name, last_name)

	// 交换两个变量的值
	var one, two = 10, 20
	one, two = two, one
	fmt.Println(one, two)
}
