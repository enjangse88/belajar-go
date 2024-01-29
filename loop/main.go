package main

import "fmt"

func main() {

	// for i:=1; i<10; i++ {
	// 	fmt.Println("Go : ",i)
	// }

	 title := "Sedang Belajar Golang"

	// for index, letter := range title {
	// 	fmt.Println("index: ", index, " letter : ", letter, "convert: ", string(letter))
	// } 

	for _, letter := range title {
		fmt.Println("letter: ", string(letter))
	}


}