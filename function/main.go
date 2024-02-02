package main

import "fmt"

func main() {

// 	sentence := printMyFunction("Saya sedang ")
// 	fmt.Println(sentence)
// result := Add(10, 5)
// fmt.Println(result)
 //  luas, keliling := Add(10,5)
 _, keliling := Add(10, 9)
 //  fmt.Println(luas)
   fmt.Println(keliling)



}

// func printMyFunction(sentence string) string {
// 	newSentence := sentence + "belajar golang"
// 	return newSentence
// }

// func Add(number, numberTwo int) int {
// 	// result := number + numberTwo
// 	// return result
// }

// func Add(number, numberTwo int) (int, int) {
// 	luas := number * numberTwo
// 	keliling := 2 * (number + numberTwo)

// 	return luas, keliling
// }

func Add(number int, numberTwo int) (luas int, keliling int){
	luas = number * numberTwo
	keliling = number * numberTwo

	return
}