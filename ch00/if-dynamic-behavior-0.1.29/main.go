package main

import "fmt"

type Crier interface {
	Cry() string
}
type Footstepper interface {
	Footsteps() string
}
type CryFootstepper interface {
	Crier
	Footstepper
}

type Person struct{}

func (p *Person) Cry() string {
	return "Hi"
}
func (p *Person) Footsteps() string {
	return "Pitapat"
}

type PartyPeople struct {
	Person
}

// Cryメソッドの実装により動的に挙動を変更
func (p *PartyPeople) Cry() string {
	return "Sup?"
}

func main() {
	var cf CryFootstepper

	cf = &Person{}
	fmt.Println(cf.Cry(), cf.Footsteps())
	// Output: Hi Pitapat

	cf = &PartyPeople{}
	fmt.Println(cf.Cry(), cf.Footsteps())
	// Output: Sup? Pitapat
}
