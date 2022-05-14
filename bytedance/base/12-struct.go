package main

import "fmt"

type user struct {
	name     string
	password string
}

func checkpass(u user, password string) bool {
	return u.password == password
}

func checkpass_ptr(u *user, password string) bool {
	return u.password == password
}

func main() {
	var user0 = user{name: "XiaoMing", password: "1234"}
	var user1 = user{"XiaoMing", "1234"}
	var user2 = user{name: "XiaoMing"}
	user2.password = "123"

	var user3 user
	user3.name = "Xiao"
	user3.password = "123"

	fmt.Println(user0, user1, user2, user3)

	fmt.Println(checkpass(user0, "1234"))
	fmt.Println(checkpass_ptr(&user0, "1234"))
}
