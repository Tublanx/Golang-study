package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"이기현", "2174", 19, 10}
	vip := VIPUser{
		User{"김신영", "2172", 19, 10},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d 유저 레벨:%d\n", vip.Name, vip.ID, vip.Age, vip.Level, vip.User.Level)
}
