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
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Nanang"
	// user.LastName = "Baranang"
	// user.Email = "varanang@nanang.id"
	// user.isActive = true

	//user := User{ID: 1, FirstName: "Jajang", LastName: "Surajang", Email:"surajang@jajang.id", isActive: true,}

	user := User{
		Email:"surajang@jajang.id",
		ID: 1,
		FirstName: "Jajang",
	    LastName: "Surajang",
	    isActive: true,}


	fmt.Println(user)

}