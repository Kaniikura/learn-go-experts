package main

import "fmt"

func main() {
	msg := "Go Expert"

	// 文字列からバイトスライスにキャスト
	bs := []byte(msg)
	fmt.Println(bs) // Output: [71 111 32 69 120 112 101 114 116]

	// バイトスライスから文字列へキャストを行う
	s := string(bs)
	fmt.Println(s) // Output: Go Expert
}
