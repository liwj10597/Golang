package main

import "fmt"

/**
  类型
*/
func main() {
	// 值类型
	var i = 1001
	var j = i
	// 获取值类型的内存地址，将i赋值给j，是将内存中的i值copy一份给j 值类型的值存储在栈中
	fmt.Print(&i, &j)

	// 引用类型
	//引用类型的变量存储的是
}
