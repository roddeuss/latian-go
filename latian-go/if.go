package main

import "fmt"

func main() {
	name := "andreasss"

	if name == "andreas" {
		fmt.Println("Hello Andreas")
	} else if name == "joko" {
		fmt.Println("Hello Joko")
	} else if name == "budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	if length := len(name); length > 7 {
		fmt.Println("Nama Anda Sangat Panjang")
	} else if length := len(name); length > 5 {
		fmt.Println("Nama Anda Cukup Panjang")
	} else {
		fmt.Println("Nama Anda sudah benar")
	}
}
