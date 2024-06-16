package main

import "fmt"

type User struct {
	ID       int
	Username string
}

func main() {
	var (
		user  User  // userはUser型とであると宣言
		pUser *User // User型を指すためのポインタ型であると宣言
	)
	addUser := &User{} // User構造体を初期化したポインタを代入

	fmt.Printf("%+v\n%+v\n%+v\n", user, pUser, addUser)
}
