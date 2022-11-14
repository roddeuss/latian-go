package main

import "fmt"

func getFullName() (string, string, string) {
	return "Andreas", "Bilmar", "Harteveld"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
