package main

import "fmt"

/**
switch语句
*/

/**
var1 可以是任意类型 val1 和val2必须是相同的类型。每个case后不用加break
*/
/*switch var1 {
case val1:
	...
case val2:
	...
default:
	...
}*/

func main() {
	char := 'B'
	switch char {
	case 'B':
		fmt.Println("B")
	case 'A':
		fmt.Println("A")
	default:
		fmt.Println("default")
	}

	grade := "优秀"
	switch {
	case grade == "优秀":
		fmt.Println("优秀")
	case grade == "良好":
		fmt.Println("良好")
	}

	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	// fallthrough
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

}
