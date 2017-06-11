package main

import (
	"fmt"
)

const (
	a = iota      // iota 常量组内第一次出现时值为 0
	b             // 不指定类型和初始值，值为上一行的右值，此时 iota 为 1
	c = iota * 10 // 此时 iota 为 2
	d
	e = 100
	f
	g = iota // 按行数递增，iota 为 6
	h
)

const (
	_, _ = iota, iota * 10 // iota 出现在每一行时递增，不是每一次出现
	i, x
	j, y
)

// units
const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func main() {
	fmt.Printf("a = %v, b = %v, c = %v, d = %v, e = %v, f = %v, g = %v, h = %v\n", a, b, c, d, e, f, g, h)
	fmt.Printf("i = %v, j = %v\n", i, j)
	fmt.Printf("x = %v, y = %v\n", x, y)
	fmt.Printf("KiB = %v, MiB = %v, GiB = %v\n", KiB, MiB, GiB)
}
