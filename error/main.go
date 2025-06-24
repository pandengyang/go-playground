package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	err := os.Remove("abc.txt")
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		fmt.Println("no such file or directory, ignore")
	}

	fmt.Println(err.Error())
	fmt.Println(err)
}
