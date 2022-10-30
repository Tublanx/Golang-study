package main

import "fmt"

func main() {
	a := 324.13455
	c := 3.14

	fmt.Printf("%08.2f\n", a)
	fmt.Printf("%08.2g\n", a)
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%f\n", c)
}
