package main

import "fmt"

func main() {
	number := 7

	switch {
	case number > 8: 
	    fmt.Println("A")
	case number > 5:
		fmt.Println("B")
	case number > 3:
		fmt.Println("C")
	default:
		fmt.Println("Default")
	}

}