package main

import "fmt"

// index 或者 key
const (
	Enone = iota
	Eio
	Einval
)

func main() {

	a := [...]string{Enone: "Enone", Eio: "Eio", Einval: "Einval"}
	s := []string{Enone: "Enone", Eio: "Eio", Einval: "Einval"}
	m := map[int]string{Enone: "Enone", Eio: "Eio", Einval: "Einval"}

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(m)
}
