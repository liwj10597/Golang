package main

import "fmt"

/**
运算符
*/
func main() {
	// 算数运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	// 关系运算符

	// 逻辑运算符
	var one bool = true
	var two bool = false
	if one && two {
		fmt.Println("one")
	}
	if one || two {
		fmt.Println("two")
	}

	// 位运算符 &  | ^  <<  >>

	// 赋值运算符

	// 其他运算符 &  *  取地址；指针变量

	// 运算符优先级

}
