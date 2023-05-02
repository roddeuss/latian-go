package main

import "fmt"

func factorialLopp(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = di kali i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLopp(15)
	fmt.Println(loop)
	fmt.Println(factorialRecursive(15))
}
