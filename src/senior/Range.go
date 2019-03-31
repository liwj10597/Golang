package main

import "fmt"

/**
range
*/
func main() {
	//这是我们使用range去求一个slice的和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// 有索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index=", i)
		}
	}

	// range 也可以用在map的键值对上
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "102" {
		fmt.Println(i, c)
	}
}
