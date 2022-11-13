package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	var b int = FloatPI * 100

	fmt.Println(a)
	fmt.Println(b)
}
