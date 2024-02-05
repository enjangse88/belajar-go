package main

import "fmt"

func main() {

	// for i:=1; i<10; i++ {
	// 	fmt.Println("Go : ",i)
	// }

	// title := "Sedang Belajar Golang"

	// for index, letter := range title {
	// 	fmt.Println("index: ", index, " letter : ", letter, "convert: ", string(letter))
	// } 

	// for _, letter := range title {
	// 	fmt.Println("letter: ", string(letter))
	// }

	title := "Golang the best language"

	//var a,i,u,e,o string

	for index, letter := range title {
		
		// if index % 2 == 0 {
		// 	fmt.Println("index: ", index, "=> ", string(letter))
		// }  

		//fmt.Println(string(letter))

		//  switch  {
		//  case string(letter) == "a": 
		//     fmt.Println("index", index, " A")
		//  case string(letter) == "i": 
		//     fmt.Println("index", index, " I")
		//  case string(letter) == "u": 
		//     fmt.Println("index", index, " U")
		//  case string(letter) == "e": 
		//     fmt.Println("index", index, " E")
		//  case string(letter) == "o": 
		//     fmt.Println("index", index, " O")
		// }

		switch string(letter) {
		case "a","i","u","e","o":
			   fmt.Println("index", index, " ", string(letter))
		}
	}




}