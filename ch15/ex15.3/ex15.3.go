package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char) // 타입 출력
	fmt.Println(char)        // char값 출력
	fmt.Printf("%c\n", char) // 문자 출력
}
