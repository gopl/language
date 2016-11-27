package main

import "fmt"

func main() {

	done := make(chan struct{})

	go func() {
		fmt.Println("go")
		close(done)
	}()

	<-done
	fmt.Println("exit")
}
