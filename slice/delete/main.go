package main

import "fmt"


func deleteString(v []string, index int) []string {
	if index >=0 && index < len(v) {
		return append(v[:index], v[index+1:]...)
	}
	return v
}

func main() {
	s := []string{"Hello1", "Hello2", "Hello3"}
	fmt.Println(s)
	s = deleteString(s, 1)
	fmt.Println(s)
	s = deleteString(s, -1)
	fmt.Println(s)
	s = deleteString(s, -2)
	fmt.Println(s)
}
