package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	io.EOF = errors.New("hello")

	buf := make([]byte, 1024)
	tempFile, _ := ioutil.TempFile("", "")
	defer os.RemoveAll(tempFile.Name())

	n, err := tempFile.Read(buf)
	fmt.Println(n, err)
}
