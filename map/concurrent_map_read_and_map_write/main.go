package main

import (
	"time"
)

func main() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"]++
			time.Sleep(time.Nanosecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Nanosecond)
		}
	}()
	select {}
}
