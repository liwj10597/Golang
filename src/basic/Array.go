package main

import "fmt"

/**
数组
*/
func main() {
	// 声明
	//var arr [3]int
	// 声明并赋值
	var arr = [3]int{1, 2, 3}

	// 取数组值
	var i = 0
	for i = 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

	// 向函数传递数组
}

/*func getAvg(num []float32) float32 {
	var i int
	var sum float32
	for i = 0; i < len(num); i++ {
		sum += num[i]
	}
	var avg float32 = sum / len(num)
	return avg
}

func getAverage(arr []int, size int) float32 {
	var i int
	var avg, sum float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = sum / size

	return avg;
}*/
