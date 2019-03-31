package main

import (
	"fmt"
	"unsafe"
)

/**
常量
*/

// 常量可以做枚举类型
const (
	ONE = 1
	TWO = 2
	THR = len("th")
	// sizeof 在编译期求值，因为可以赋值给常量
	FOUR = unsafe.Sizeof(ONE)
)

const (
	//iota = 1
	// 第一个iota = 0 每一行加1，当它在新的一行被使用时，它当前的值就会被赋给当前行
	zero   = iota
	first  = iota
	second = iota
	thrid  = iota
	fourth = 4
	impfourth
	sixth = iota
	senventh
)

// 常量
func main() {
	// 常量只能是bool 数字 字符串类型。常量声明时必须初始化
	const name string = "aubrey"
	// 常量不允许修改
	//name = ""
	fmt.Println(name)

	// 一行 声明多个常量
	const corp, address = "", ""

	fmt.Println(ONE)
	fmt.Println(THR)
	fmt.Println(FOUR)

	fmt.Println(zero)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(thrid)
	fmt.Println(fourth)
	fmt.Println(impfourth)
	fmt.Println(sixth)
	fmt.Println(senventh)
}
