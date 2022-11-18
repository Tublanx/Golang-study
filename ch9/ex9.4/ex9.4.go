package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // false로 인해 쇼트서킷이 발생하여 IncreaseAndReturn() 실행 x
		fmt.Println("1 증가")
	}

	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}

	fmt.Println("cnt: ", cnt)
}
