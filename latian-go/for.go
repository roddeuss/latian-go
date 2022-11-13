package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke", counter)
	}

	slice := []string{"Andreas", "Bilmar", "Harteveld"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index ke ", i, "=", value)
	}

	person := make(map[string]string)
	person["Name"] = "Andreas"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
