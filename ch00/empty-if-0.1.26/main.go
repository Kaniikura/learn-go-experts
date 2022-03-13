package main

import "fmt"

func main() {
	ss := []string{"John", "Richard"}

	/*
		var i interface{} = ss // OK
		var is []interface{} = ss // NG
	*/

	// スライスに代入するには繰り返しで行う
	is := make([]interface{}, 0, len(ss))
	for _, s := range ss {
		is = append(is, s)
	}
	fmt.Printf("%v\n", is)
}
