/*
检索map的某个key时,
如果只用一个变量接收,单key不存在时返回value的零值
如果用两个变量接收,第二个变量(bool类型)用于指示key是否存在
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	i := m["three"]
	fmt.Println(`m["three"] =`, i)

	if j, ok := m["four"]; ok {
		fmt.Println(`m["four"] =`, j)
	} else {
		fmt.Println(`m["four"] is null`)
	}
}
