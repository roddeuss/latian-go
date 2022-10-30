package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	type Job bool

	var noKTPEko NoKTP = "12712313131"
	var status Married = false
	var Pekerjaan Job = true

	fmt.Println(noKTPEko)
	fmt.Println(status)
	fmt.Println(Pekerjaan)
}
