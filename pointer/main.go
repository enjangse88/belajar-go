package main

import "fmt"

func main() {

	// numberA := 5
	// numberB := &numberA //proses referensing

	// fmt.Println(numberA)
	// fmt.Println(numberB) // memory point 
	// fmt.Println(*numberB) // proses deferensing
	number := 5

	fmt.Println("alamat memory",&number)
	fmt.Println("nilai awal:",number)

	change(&number, 100)
	fmt.Println("nilai akhir",number)
	fmt.Println("alamat memory",&number)

}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory", old)
	fmt.Println("Di dalam function", *old)

}