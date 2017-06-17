package main

import (
	"fmt"
)

func main() {
	a := [10]struct{}{}
	b := a[:]
	c := [0]int{}
	d := struct{}{}
	fmt.Printf("% p, %p, %p, %p\n", &a[0], &b[0], &c, &d)
}
