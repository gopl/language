package main

import "fmt"

type foo struct {
	s string
}

type bar struct {
	s int
}

type fooBar struct {
	foo
	bar
}

func main() {
	var fb = fooBar{}
	fmt.Printf("%#v %#v %#v\n", fb, fb.foo.s, fb.bar.s)
	// fmt.Println(fb.s) // ambiguous selector fb.s
}
