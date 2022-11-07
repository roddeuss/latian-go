package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andreas"
	names[1] = "Xavier"
	names[2] = "Greedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		100,
		99,
		98,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}
