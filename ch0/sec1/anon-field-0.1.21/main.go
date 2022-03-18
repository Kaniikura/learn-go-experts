package main

import "fmt"

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func main() {
	c := Card{
		string: "Credit",
		Chip: Chip{
			Number: 42_424_242_424_242,
		},

		Number: 54_545_454_545_454,
	}

	// CardにはScanメソッドがないため、c.Chip.Scan()が実行される
	c.Scan()
	// Scanメソッドのレシーバーは、CardではなくChipであることがわかる
	// Output: 4242424242424242
}
