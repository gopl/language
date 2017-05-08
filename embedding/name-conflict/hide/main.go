package main

import "fmt"

type foo struct {
	s string
}

type bar struct {
	foo
}

type bar2 struct {
	foo
	s int
}

func main() {
	var b = bar{foo{"a"}}
	fmt.Println(b.s)

	var b2 =bar2{s:1}
	fmt.Println(b2.s)
}
