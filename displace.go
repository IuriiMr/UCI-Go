package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5 * a * t * t + v0 * t + s0
	}
	return fn
}

func main() {

	var a, v0, s0, t float64

	fmt.Println("Please, enter a value for acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Println("Please, enter a value for initial velocity: ")
	fmt.Scanf("%f", &v0)
	fmt.Println("Please, enter a value for initial displacement: ")
	fmt.Scanf("%f", &s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Displacement will be computed after the entered time")
	fmt.Println("(press ENTER after each value or just ENTER when done)")

	for {
		fmt.Println("Please, enter a value for time: ")
		_, err := fmt.Scanf("%f", &t)
		if err == nil {
			fmt.Println("The displacement is:", fn(t))
		} else {
			break
		}
	}
}
