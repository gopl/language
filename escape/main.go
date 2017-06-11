package main

// go build -gcflags "-l -m"

import "fmt"

func test() *int {
	a := 0x100
	return &a
}
func main() {
	a := test()
	fmt.Println(a, *a)
}
