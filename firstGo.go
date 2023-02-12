package firstGo

import (
	"fmt"
)

type User struct {
	nickname string
	email    string
}

type Recipe struct {
	name    string
	calorie int
	author  string
	cat     string
}

func UserInfo() {
	u := User{"Nurly", "nurly125@mail.ru"}
	fmt.Printf("New user nickname is %s and email - %s\n", u.nickname, u.email)
}

func RecipeInfo() {
	r := Recipe{"Napoleon", 406, "Nurly", "dessert, cake"}
	fmt.Printf("New recipe is %s - %s, it's calorie - %d, and Author is %s\n", r.cat, r.name, r.calorie, r.author)
}

func AuthorCheck() {
	u := User{"Nurly", "nurly125@mail.ru"}
	r := Recipe{"Napoleon", 406, "Nurly", "dessert, cake"}

	if u.nickname == r.author {
		fmt.Printf("New User %s is author of %s\n", u.nickname, r.name)
	}
}

func CalorieCheck() {
	r := Recipe{"Napoleon", 406, "Nurly", "dessert, cake"}

	if r.calorie > 400 {
		fmt.Printf("New %s's calorie is %d, it is bigger than normal", r.name, r.calorie)
	}
}
