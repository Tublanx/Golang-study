package main

import "fmt"

func main() {
	var (
		a, b int = 10, 20
	)

	a, b = b, a

	fmt.Println(a, b)
}
