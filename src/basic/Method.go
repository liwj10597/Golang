package main

import "fmt"

/**
方法：方法在func关键字后是接收者而不是函数名，接收者可以是自己定义的一个类型，这个类型可以是struct，interface
*/

/**
  定义结构体
*/
type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 4
	fmt.Println(c.getArea())
}

// 该method属于Circle 类型对象的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
