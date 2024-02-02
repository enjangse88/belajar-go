package main

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	isActive bool
}

type Group struct {
	Name  string
	Admin  User
	Users []User
	isAvaliable bool
}

func main() {
	user1 := User{1, "Zelda", "Nintendo", "zelda@nintentdo.com", true}
	user2 := User{2, "Mario", "Nintendo", "mario@nintendo.com", true}
	// // displayUser1 := displayUser(user1)
	// // displayUser2 := displayUser(user2)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)
	users := []User{user1, user2}

	group := Group{"Game", user1, users, true}
	displayGroup(group)

}

func displayGroup(group Group) {
	fmt.Printf("Name : %s ", group.Name )
	fmt.Println("")
	fmt.Printf("Member Count: %d ", len(group.Users) )
	fmt.Println("")

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}