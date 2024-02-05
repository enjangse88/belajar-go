package main

import "fmt"

func main() {

	// Map_Name := map[Data_type]Map_type{"key":"value", "key":"value"}
	// var myMap map[string]int
	// myMap = map[string]int{}
	// myMap["ruby"] = 9
	// myMap["JavaScript"] = 8
	// myMap["Laravel"] = 8

	//myMap := map[string]int{"ruby":8, "golang":9}
	myMap := map[string]string{
		"ruby":"beautiful",
	   "golang":"super fast",
	   "Python":"Flexible",
	}

	for key, value := range myMap {
		fmt.Println("Key : ",key," Value : ", value)
	}

	fmt.Println("===========")
	delete(myMap, "golang")

	for key, value := range myMap {
		fmt.Println("Key : ",key," Value : ", value)
	}

	value := myMap["ruby"]
	fmt.Println(value)

	isAvaliable := myMap["JavaScript"]
	fmt.Println(isAvaliable)




	//fmt.Println(myMap)
}