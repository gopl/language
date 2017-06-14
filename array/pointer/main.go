package main

import "fmt"

func main() {
	a := [...]int{1, 2}
	fmt.Printf("%T %p %v\n", &a, &a, &a)
	fmt.Printf("%T %p %v\n", &a[0], &a[0], &a[0])
	fmt.Printf("%T %p %v\n", &a[1], &a[1], &a[1])
}
