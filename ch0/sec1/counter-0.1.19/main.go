package main

import (
	"fmt"

	"example.com/m/v1/syncutil"
)

func main() {
	c := &syncutil.Counter{
		Name: "Access",
	}

	fmt.Println(c.Increment())
	fmt.Println(c.View())
}
