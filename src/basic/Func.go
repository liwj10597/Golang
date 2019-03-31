package main

import (
	"fmt"
	"math"
)

// 初始化方法每个包可以有多个，自动调用，且只调用一次
func init() {
	fmt.Println("call init")
}

/**
函数
*/
func main() {
	// 简单函数
	fmt.Println(Max(3, 2))

	// 返回多个值的函数 采用值传递
	var x, y = "100", "200"
	fmt.Println(Swap(x, y))
	fmt.Println(x, y)
	// 函数传值 值传递 引用传递 默认情况下，go使用的是值传递

	fmt.Println(SwapRef(&x, &y))
	fmt.Println(x, y)

	fmt.Println(SwapRefTmp(&x, &y))
	fmt.Println(x, y)

	// 函数作为值使用
	getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquare(9))

	// 函数闭包 go语言支持匿名函数，可作为闭包
	nextNumber := getNextNumber()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

}

/**
使用函数闭包 。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。返回的是一个方法
*/
func getNextNumber() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

/**
获取两个整数的最大值
*/
func Max(numOne, numTwo int) int {
	var result int
	if numOne > numTwo {
		result = numOne
	} else {
		result = numTwo
	}
	return result
}

/**
函数返回多个值 值传递
*/
func Swap(x, y string) (string, string) {
	return y, x
}

/**
引用传递一：传递了引用，仅仅是交换了引用对应的值，并没有改变引用对应的值
*/
func SwapRef(x, y *string) (string, string) {
	return *y, *x
}

/**
引用传递二：传递了引用，交换了引用对应的值，交换后，影响实参
*/
func SwapRefTmp(x, y *string) (string, string) {
	var tmp string
	tmp = *x
	*x = *y
	*y = tmp
	return *x, *y
}
