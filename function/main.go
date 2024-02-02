package main

import "fmt"


func main() {

// 	sentence := printMyFunction("Saya sedang ")
// 	fmt.Println(sentence)
result := Add(10, 5)
fmt.Println(result)


}

// func printMyFunction(sentence string) string {
// 	newSentence := sentence + "belajar golang"
// 	return newSentence
// }

func Add(number, numberTwo int) int {
	result := number + numberTwo
	return result
}
