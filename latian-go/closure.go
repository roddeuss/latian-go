package main

import "fmt"

func main() {
	name := "Budi"
	counter := 0

	increment := func() {
		name := "Andreas"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
