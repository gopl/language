package main

import (
	"fmt"

	log1 "github.com/gostudying/language/same_package_name/log1"
	log2 "github.com/gostudying/language/same_package_name/log2"
)

func main() {
	fmt.Println(log1.Logger)
	fmt.Println(log2.Logger)
}
