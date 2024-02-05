package main

import "fmt"

func main() {

	//var gameConsoles [6]string <=== Array
	//var gameConsoles []string <== slice

	var gameConsoles []string

	gameConsoles = append(gameConsoles, "Playstation")
	gameConsoles = append(gameConsoles, "XBOX")
	gameConsoles = append(gameConsoles, "Nintendo GameCube")

	for _, console := range gameConsoles {
		fmt.Println(console)
	}

	fmt.Println(gameConsoles)
}