package main

import "fmt"

// S ...
type S struct {
	i int
	s string
}

// T ...
type T struct {
	i int
	s string
}

func main() {
	t := T{0, ""}
	s := S(t)
	fmt.Printf("%#v", s)
}
