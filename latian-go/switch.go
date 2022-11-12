package main

import "fmt"

func main() {

	name := "Andreas"

	switch name {
	case "Andreas":
		fmt.Println("Hallo Andreas")
		fmt.Println("Hello Andreas")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Dong")
		fmt.Println("Kenalan Dong")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama Terlalu Panjang")
	//case false:
	//	fmt.Println("Nama Anda Benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
