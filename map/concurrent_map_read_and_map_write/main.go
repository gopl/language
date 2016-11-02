package main

import (
	"fmt"
	"sync"
)

func main() {
	const (
		R = 1000
		W = 3
	)
	wg := sync.WaitGroup{}
	m := map[int]int{}

	wg.Add(R + W)
	for i := 0; i < R; i++ {
		go func() {
			fmt.Println(m)
			wg.Done()
		}()
	}
	for i := 0; i < W; i++ {
		go func() {
			m[i%13] = i
			wg.Done()
		}()
	}

	wg.Wait()
}
