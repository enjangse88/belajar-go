package main

import (
	"fmt"
	"errors"
)

func main() {

	// scores := []int{10, 5, 8, 7,9}
	// total := sum(scores)

	// fmt.Println(total)

	result, err := calculator(5, 8, "=")
	//result := calculator(5, 8, "*")
	if err != nil{
	   fmt.Println("Terjadi kesalahan", err)
	} else {
		fmt.Println(result)
	}
	
	

}

// func sum(numbers  []int) int{
// 	var total int
	
// 	for _, number := range numbers{
// 		total = total + number
// 	}
// 	return total
// }

func calculator(number int, numberTwo int, operator string) (int, error){

	var result int
	var errorResult error

	switch operator {
	case "+": result = number + numberTwo
	case "-": result = number - numberTwo
	case "*": result = number * numberTwo
	case "/": result = number / numberTwo
	default: errorResult =  errors.New("Unknow Operator")
	}

	return result, errorResult 
}