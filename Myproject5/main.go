package main

import (
	"fmt"
	"myproject/mathutils"
)

func main() {
	fmt.Println("Programma kosyldy")
	result := mathutils.Add(14, 2)
	fmt.Println("14 + 2 =", result)
}
