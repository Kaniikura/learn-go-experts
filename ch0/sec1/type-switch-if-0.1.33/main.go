package main

import "fmt"

func main() {
	i := interface{}("Go Expert")

	// 型情報を取り出して、それぞれの型情報ごとに処理をわける
	switch i.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("This is integer,", i)
	case string:
		fmt.Println("This is string,", i)
	default:
		fmt.Printf("This is unknown type, %T\n", i)
	}
}
