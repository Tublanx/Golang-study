package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // 매개변수로 Data 포인터를 받는다.
	arg.value = 999 // arg 데이터 변경
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data) // 인수로 data의 주소를 넘긴다.
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
