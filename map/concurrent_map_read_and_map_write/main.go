package main

import (
	"time"
)

func main() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"] += 1
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
