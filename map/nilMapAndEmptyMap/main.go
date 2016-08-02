// map必须通过make创建后,才能进行赋值操作

package main

import "fmt"

func main() {
	var m1 map[string]int
	m2 := make(map[string]int)

	fmt.Println("m1 =", m1)
	fmt.Println("m2 =", m2)

	m1["m1"] = 1 // panic: assignment to entry in nil map
	m2["m2"] = 2

}
