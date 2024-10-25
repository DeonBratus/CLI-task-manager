package tskMng

import "fmt"

type User struct {
	userID   int    `json:"UserID"`
	UserName string `json:"UserName"`
}

func NewUser(id int, name string) User {
	newUser := User{id, name}
	fmt.Println("New user has been created")
	return newUser
}
