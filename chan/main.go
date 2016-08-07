package main

import "fmt"

func main() {
	ch := make(chan string)
	toMain := make(chan bool)
	go receiveOnly(toMain, ch)
	go sendOnly(toMain, ch)
	<-toMain
	<-toMain
}

// 只写channel chan<-
func sendOnly(toMain chan<- bool, ch chan<- string) {
	ch <- "This is a secrete."
	fmt.Println("Send a secrete")
	toMain <- true
}

// 只读channel <-chan
func receiveOnly(toMain chan<- bool, ch <-chan string) {
	msg := <-ch
	fmt.Println("Receive:", msg)
	toMain <- true
}
