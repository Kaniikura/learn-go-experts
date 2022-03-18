package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Aging() {
	u.Age++
}

func (u User) AgingButNothingHappen() {
	u.Age++
}

func main() {
	u := &User{
		Name: "Richard",
		Age:  33,
	}

	// レシーバがUserのポインタなので、呼び出し元のuと同一のuを操作できる
	u.Aging()
	fmt.Println(u.Age) // Output: 34

	// ポインタ型であっても、型がレシーバのメソッドは呼び出すことができる
	u.AgingButNothingHappen()
	// レシーバーがUserなので、呼び出し元には変更が波及しない
	fmt.Println(u.Age) // Output: 34
}
