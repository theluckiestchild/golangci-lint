package main

import (
	"fmt"
	"os"
)

var a int

func main() {
	a = 1
	c := os.Args
	fmt.Println(c)
}
