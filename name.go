package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}
func (Cow) Eat() {
	fmt.Println("grass")
}
func (Cow) Move() {
	fmt.Println("walk")
}
func (Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {}
func (Bird) Eat() {
fmt.Println("worms")
}
func (Bird) Move() {
	fmt.Println("fly")
}
func (Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {}
func (Snake) Eat() {
	fmt.Println("mice")
}
func (Snake) Move() {
	fmt.Println("slither")
}
func (Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	m := map[string]Animal{}
	for {
		fmt.Println(">")
		var act, name, inf string
		fmt.Scan(&act, &name, &inf)

		switch act {
		case "newanimal":
			switch inf {
			case "cow":
				m[name] = Cow{}
			case "bird":
				m[name] = Bird{}
			case "snake":
				m[name] = Snake{}
			}

			fmt.Println("Created it!")
		case "query":
			switch inf {
			case "eat":
				m[name].Eat()
			case "move":
				m[name].Move()
			case "speak":
				m[name].Speak()
			}
		}
	}
}
