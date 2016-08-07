package main

import "fmt"

func main() {

	type pi *int
	type pj *int

	i := 1
	var p1 pi = &i
	fmt.Println("pi =", p1)

	var p2 pj = pj(p1)
	fmt.Println("pj =", p2)

}
