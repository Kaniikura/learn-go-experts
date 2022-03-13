package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// メソッド(func() string)を値として扱う
	fv := Hex(1024).String
	fmt.Println(fv()) // Output: 400

	// メソッド(func(Hex) string)を式として扱う
	fe := Hex.String
	fmt.Println(fe(1024)) // Output: 400
}
