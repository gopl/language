package main

import "fmt"

func main() {

	// 可以是编译器能计算出结果的表达式
	const length = len("hello, world")
	fmt.Printf("length = %v\n", length)

	// 常量组内，不指定类型和初始值时与上一行非空常量右值相同
	const (
		i = 100
		j

		x = "abc"
		y
		z
	)
	fmt.Printf("i = %v, j = %v\n", i, j)
	fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
}
