package main

import "fmt"

type ErrNoSuchEntity struct{ error }
type ErrConflictEntity struct{ error }

func main() {
	do := func() error {
		return &ErrConflictEntity{}
	}

	switch do().(type) {
	case nil:
		// do nothing
	case *ErrNoSuchEntity:
		fmt.Println("error no such entity")
	case *ErrConflictEntity:
		fmt.Println("error conflict entity")
	default:
		fmt.Print("unknown error")
	}

	// Output: error conflict entity
}
