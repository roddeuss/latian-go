package main

import "fmt"

func main() {
	var name string

	name = "Andreas Bilmar "
	fmt.Println(name)

	name = "Andreas Girsang"
	fmt.Println(name)

	var friendName = "Bilmar"
	fmt.Println(friendName)

	var age = 23
	fmt.Println(age)

	country := "indonesia"
	fmt.Println(country)

	var (
		firstName = "Andreas"
		lastName  = "Bilmar"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + lastName)
}
