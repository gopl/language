package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a =", a)

	s := a[3:5:6] // 指定 cap，限定对底层数组能够 append 范围
	fmt.Printf("s = %v\t\t(%#v)\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))

	s = append(s, 0)
	fmt.Printf("s = %v\t\t(%#v)\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	fmt.Println("a =", a)

	s = append(s, 0)
	fmt.Printf("s = %v\t\t(%#v)\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s))) // 重新分配内存
	fmt.Println("a =", a)
}
