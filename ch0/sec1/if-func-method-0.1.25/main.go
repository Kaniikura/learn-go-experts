package main

import "fmt"

type Crier interface {
	Cry() string
}

type ParrotFunc func() string

func (p ParrotFunc) Cry() string {
	return p()
}

func main() {
	// 型リテラルで無名関数を定義し、ParrotFunc型にキャスト
	var c Crier = ParrotFunc(func() string {
		return "Squawk"
	})

	fmt.Println(c.Cry()) // Output: Squawk
}
