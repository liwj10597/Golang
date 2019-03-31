package main

import "fmt"

/**
空白标识符
*/
func main() {
	// 声明了twoParam 但是却不想使用，为了避免无法编译的情况出现,使用空白符接收函数的返回值并舍弃掉
	var oneParam, twoParam int
	var threeParam float32
	oneParam, _, threeParam = TestBlankSymbol()
	fmt.Printf("oneParam=%d, twoParam=%d, threeParam=%d", oneParam, twoParam, threeParam)
}

// 返回多个值
func TestBlankSymbol() (int, int, float32) {
	return 5, 7, 9.5
}
