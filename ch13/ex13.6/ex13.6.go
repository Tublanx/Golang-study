package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{19, 88.3}
	fmt.Println(unsafe.Sizeof(user))
}
