package main

import "fmt"

func main() {
	type file struct {
		name string
		attr struct { // 匿名结构作为字段
			owner int
			perm  int
		}
	}
	f := file{
		name: "name",
		attr: struct { // 再次定义匿名结构
			owner int
			perm  int
		}{
			1,
			0755,
		},
	}

	f.attr.owner = 1 // 单独赋值
	f.attr.perm = 0755

	fmt.Println(f)
}
