package main

import "fmt"

func main() {
	a := [...]int{1, 2}
	pa := &a
	pa[1] += 100
	fmt.Println("a:", a)

	s := []int{3, 4}
	ps := &s
	// ps[1] += 100 -- invalid operation: ps[1] (type *[]int does not support indexing)
	(*ps)[1] += 100 // 先转换成 []int
	fmt.Println("s:", s)
}
