package main

import "fmt"

func main() {
	// 큰따음표로 묶으면 특수 문자가 동작
	str1 := "Hello\t'World'\n"

	// 백쿼트로 묶으면 특수 문자가 동작하지 않음
	str2 := `Go is "awesome"!\nGo is simple and\t 'powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}
