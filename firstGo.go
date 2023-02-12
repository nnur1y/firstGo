package firstGo

import (
	"fmt"
)

type User struct {
	nickname string
	email    string
}

type Recipe struct {
	name     string
	calorie  int
	author   string
	category string
}

func UserInfo(nick string, email string) *User {
	u := User{nick, email}
	return &u
}

func (user User) newUser() {
	fmt.Println("Nickname:", user.nickname, "\nEmail:", user.email)
}

func RecipeInfo(name string, calorie int, author string, category string) *Recipe {
	r := Recipe{name, calorie, author, category}
	return &r
}
