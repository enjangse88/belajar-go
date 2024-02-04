package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Final Fantasy IX")
	gamer.AddGame("Suikoden")
	gamer.AddGame("Syphon filter")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}

