package main

import "fmt"

func getMyAge() (int, bool) {
	return 19, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("you are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("you are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("your age is", age) 참조 불가능 (age 소멸)
}
