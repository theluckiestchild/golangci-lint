package main

import (
	"fmt"
	"os"
)

var a int

func main() {
	a = 10
	c := os.Args
	fmt.Println(c)
}
