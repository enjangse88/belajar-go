package main

import "fmt"

func main() {

	//countries := [6]string {"Arab", "Japan", "Italia", "Indonesia", "USA", "Netherlands"}
    countries := [...]string {
		"Arab",
		"Japan",
		"Italia",
		"Indonesia",
		"USA",
		"Netherlands"}

	for index, country := range countries {
		fmt.Println("index: ", index, "name : ", country)
	}
	fmt.Println(countries)
}