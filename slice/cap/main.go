package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a =", a)

	s1 := a[3:7]
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n", s1, len(s1), cap(s1))

	s2 := a[3:5:7] // 限定切片 cap
	fmt.Printf("s2 = %v, len(s2) = %v, cap(s2) = %v\n", s2, len(s2), cap(s2))
}
