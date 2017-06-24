package main

import (
	"fmt"
	"reflect"
)

// S ...
type S struct {
}

// T ...
type T struct {
	S
}

// SVal ...
func (S) SVal() {}

// SPtr ...
func (*S) SPtr() {}

// TVal ...
func (T) TVal() {}

// TPtr ...
func (*T) TPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T
	methodSet(t)
	fmt.Println("------")
	methodSet(&t)
}
