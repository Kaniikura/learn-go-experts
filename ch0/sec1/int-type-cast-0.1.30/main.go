package main

import "fmt"

func main() {
	var i int32 = 100
	var j int64

	// 暗黙的なキャストがないので、コンパイル時にエラーになる
	// j = i

	// 明示的にキャストすることで代入できる
	j = int64(i)
	fmt.Println(j) // Output: 100
}
