package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네 ㅎㅎ")
		} else {
			fmt.Println("더치페이 ㄱㄱ")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이 ㅎ..")
		} else {
			fmt.Println("사람도 얼마 없는데 더치페이 ㄱㄱ")
		}
	} else {
		fmt.Println("오늘 내가 골든벨 울린다 ㅋㅋ")
	}
}
