package main

import (
	"fmt"
)


func main() {
	helloworld()
}

func helloworld() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
