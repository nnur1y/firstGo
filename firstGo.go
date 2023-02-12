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

func UserInfo() {
	u := User{"Nurly", "nurly125@mail.ru"}
	fmt.Printf("New user nickname is %s and email - %s", u.nickname, u.email)
}

func RecipeInfo(name string, calorie int, author string, category string) *Recipe {
	r := Recipe{name, calorie, author, category}
	return &r
}
