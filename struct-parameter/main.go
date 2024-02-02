package main

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	isActive bool
}

func main() {
	user1 := User{1, "Zelda", "Nintendo", "zelda@nintentdo.com", true}
	user2 := User{2, "Mario", "Nintendo", "mario@nintendo.com", true}
	displayUser1 := displayUser(user1)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

}

func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}