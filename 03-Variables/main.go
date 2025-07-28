package main

import "fmt"

func main() {
	//declaring variables hard-way
	var quantity int = 10
	var username string = "John"
	var isPermission bool = true
	var costPerMsg float32 = 2.89

	fmt.Println(quantity, username, isPermission, costPerMsg)

	//declaring variables easy-way
	// := short assignment operator. allows us to declare variables and go infer its type.

	//var username string.
	//is same way as:
	//string := ""

	numCars := 10 //inferred as integer.
	congrats := "Happy Birthday!"

	fmt.Println(numCars, congrats)

	costPerText := 2.0
	fmt.Printf("%T\n", costPerText)

	//Same line declaration
	mileage, company := 64727, "Tesla"

	//is the same as
	//mileage := 64727
	//company := "Tesla"

	fmt.Println(mileage, company)

	//ASSIGNMENT
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)
}
