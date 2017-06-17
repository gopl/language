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
	}

	f.attr.owner = 1 // 无法直接初始化
	f.attr.perm = 0755

	fmt.Println(f)
}
