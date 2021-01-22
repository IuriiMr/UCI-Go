package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}
func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	m := map[string] Animal {
		"cow" : Animal{"grass", "walk", "moo"},
		"bird" : Animal{"worms", "fly",  "peep"},
		"snake" : Animal{"mice", "slither", "hsss"},
	}

	fmt.Println("Please, type a request after prompt.\n" +
		"Every request must be a single line containing 2 strings.\n" +
		"The first string is the name of an animal, either “cow”, “bird”, or “snake”.\n" +
		"The second string is the name of the information requested about the animal," +
		" either “eat”, “move”, or “speak”.")

	for {

		fmt.Println(">")
		var anm, act string
		fmt.Scan(&anm, & act)

		switch act {
		case "eat":
			m[anm].Eat()
		case "move":
			m[anm].Move()
		case "speak":
			m[anm].Speak()
		default:
			fmt.Println("Unknown action! Try again!")
		}
	}
}
