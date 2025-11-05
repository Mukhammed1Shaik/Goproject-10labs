package main

import (
	"fmt"
	"myproject/calc"
	"myproject/stringutils"
)

func main() {
	text := "GoLang"
	reversed := stringutils.Reverse(text)
	squared := calc.Square(5)

	fmt.Println("Audarylgan text:", reversed)
	fmt.Println("5 kvadraty", squared)
}
