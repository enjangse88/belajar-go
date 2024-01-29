package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo belajar golang")

	result := calculation.Add(9, 4)
    
	fmt.Println(result)

	result2 := calculation.Multiple(5, 7)
	fmt.Println(result2)

	fmt.Println(calculation.Multiple(6, 6))
}
