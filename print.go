package main

import "fmt"

func printData(data interface{}) {
	ID := ""
	s := fmt.Sprintf("data %v", data)
	fmt.Print(s)
	fmt.Println(ID)
}
