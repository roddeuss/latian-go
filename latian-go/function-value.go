package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	sayGoodBye := getGoodBye

	result := sayGoodBye("Andreas")
	fmt.Println(result)
	fmt.Println(getGoodBye("Andreas"))
}
