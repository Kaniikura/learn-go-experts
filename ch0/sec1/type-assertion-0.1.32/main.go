package main

import "fmt"

func main() {
	i := interface{}("Go Expert")

	// 型アサーションを利用して、文字列として値を受け取る
	s := i.(string)
	fmt.Println(s) // Output: "Go Expert"

	// 型アサーションを利用してバイトスライスとして値を取り出そうとするが、
	// バイトスライスではないためパニックが起こってしまう
	// n := i.([]byte)
	// panic: interface conversion: interface {} is string, not [] uint8

	// 返却値を2つ受け取り、型の判定結果を受け取ることでパニックを防ぐことができる
	n, ok := i.([]byte)
	fmt.Println(n, ok) // Output: [] false
}
