package main

import "fmt"

func sayHelo() {
	fmt.Println("Hello")
}
func main() {
	for i := 0; i < 10; i++ {
		sayHelo()
	}
}
