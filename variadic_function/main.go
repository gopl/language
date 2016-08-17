package main

import "fmt"

func print(str1 string, str2 string, v ...interface{}){
	fmt.Printf("%v|%v|%v", str1, str2, fmt.Sprintln(v...))
}

func main() {
	print("hello", "world", []int{1, 2}, "hi")
}
