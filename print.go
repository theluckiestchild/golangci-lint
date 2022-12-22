package main

import "fmt"

func PrintData(data interface{}) {
	ID := ""
	s := fmt.Sprintf("data %v", data)
	fmt.Print(s)
	fmt.Println(ID)
}
