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

func newRecipe(name string, calorie int, author string, category string) *Recipe {
	r := Recipe{name, calorie, author, category}
	return &r
}

func RecipeInfo() {
	r := newRecipe("Napoleon", 406, "Nurly", "dessertcake")
	fmt.Printf("New recipe is %s - %s, it's calorie is %d, \nAuthor - %s", r.category, r.name, r.calorie, r.author)
}
