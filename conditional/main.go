package main

import "fmt"

func main() {
	Nilai := 7
	var grade string

	if Nilai > 7 {
	  grade = "B"
	} else if Nilai > 3 {
	  grade = "C"
	 } else {
		fmt.Println("Tidak lulus")
	 }
 
	fmt.Println(grade)


}