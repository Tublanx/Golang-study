package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Float).SetString("0.1") // Float 객체 생성. a는 0.1, b는 0.2, c는 0.3을 나타냄
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b) // a와 b값을 더한 값을 d에 저장
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // c와 d값을 비교. 형식은 x.Cmp(y)로 표현한다. 반환값 -1은 x가 작은 경우, 1은 x가 큰 경우, 0은 두 값이 같을 경우
}
