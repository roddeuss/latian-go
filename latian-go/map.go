package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Andreas",
		"address": "Medan",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)

	book["title"] = "belajar Go-Lang"
	book["authir"] = "Dre"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
