package main

import "fmt"

/**
切片
*/
func main() {
	// 使用make定义切片
	var slice []int = make([]int, 10)
	//或者 slice := make([]int, 10)
	fmt.Println(slice[1])

	// 创建切片并指定容量(容量不能小于长度)
	var capacitySlice = make([]int, 10, 11)
	// 可以访问len最后一个记录
	fmt.Println(capacitySlice[9])
	// 但是不可以访问超过len之后，capacity之内无初始化的记录
	//println(capacitySlice[10])

	// 切片的len和cap函数
	fmt.Println(len(capacitySlice), cap(capacitySlice), capacitySlice)

	// 切片作为函数参数
	printSlice(capacitySlice)

	// 空切片为nil
	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("切片nilSlice是空的")
	}

	// 切片截取
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
	/* 从以上例子可以看出cap的容量从startIndex到最后 */

	// copy切片
	var number = []int{-1}
	printSlice(number)

	// 允许追加空切片
	number = append(number, 0, 3, 4)
	printSlice(number)

	// 目的地切片需要通过make指定一定的容量才能容得下number切片的copy
	var des = make([]int, len(number), cap(number))
	copy(des, number)
	printSlice(des)

}

// 切片作为函数参数
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
