package main

import "fmt"

func test(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

func main() {
	a := []int{1, 2, 3}
	test(a...)
	fmt.Println(a)

	i, j, k := 10, 20, 30
	test(i, j, k)
	fmt.Println(i, j, k)
}
