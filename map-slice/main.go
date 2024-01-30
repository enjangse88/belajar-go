package main

import "fmt"

func main() {

	students := []map[string]string{
		{"name":"Agung", "score":"A"},
		{"name":"Link","score":"B"},
		{"name":"Mario","score":"C"},
	}

	for _, student := range students {
		fmt.Println(student)
	}

	fmt.Println("=============")

	for _, student := range students {
		fmt.Println("Name : ", student["name"], " Score :",student["score"])
	}


}