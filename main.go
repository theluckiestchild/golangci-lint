package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	c := os.Args
	fmt.Println(c)
}
