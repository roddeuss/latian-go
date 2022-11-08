package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktoboer",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "diubah"
	//fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSLice := make([]string, 2, 5)

	newSLice[0] = "Andreas"
	newSLice[1] = "Bilmar"
	fmt.Println(newSLice)

	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

	copySlice := make([]string, len(newSLice), cap(newSLice))
	copy(copySlice, newSLice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
