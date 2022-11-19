package main

import "fmt"

func getMyage() int {
	return 19
}

func main() {
	switch age := getMyage(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
