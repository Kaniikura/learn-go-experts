package main

import "fmt"

type Crier interface {
	// Cryメソッドを満たしていれば、Crierインタフェースを満たしていると言える
	Cry() string
}

type Duck struct{}

func (d Duck) Cry() string {
	return "Quack"
}

func main() {
	// Duck型はCryメソッドを実装しているので、Crierインタフェースを満たしている
	var c Crier = Duck{}

	fmt.Println(c.Cry())
}
