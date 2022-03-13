package main

import (
	"encoding/json"
	"fmt"
	"os"

	"example.com/m/v1/vault"
)

func main() {
	s := vault.NewSecret()

	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("failed to json encode, error =", err)
	}
}
