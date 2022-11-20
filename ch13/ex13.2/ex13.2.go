package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"이기현", "2174", 19}
	vip := VIPUser{
		User{"김신영", "2172", 19},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}
