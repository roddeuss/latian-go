package main

import "fmt"

type blackList func(string) bool

func registerUser(name string, blacklist blackList) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blaclList1(name string) bool {
//	return name == "admin"
//}
//
//func blaclListRoot(name string) bool {
//	return name == "root"
//}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("andreas", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("andreas", func(name string) bool {
		return name == "root"
	})
}
