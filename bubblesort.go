package main

import (
	"fmt"
	"strings"
)

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int) {

	length := len(sli)

	for i := 0; i < length; i++ {

		for j := 0; j < length - i - 1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Input(sli []int, err error) []int {
	if err != nil || len(sli) == 10 {
		return sli
	}
	var new int
	n, err := fmt.Scanf("%d", &new)
	if n == 1 {
		sli = append(sli, new)
	}
	return Input(sli, err)
}

func main() {

	fmt.Println("Please, type in a sequence of up to 10 integers: ")
	fmt.Println("(press ENTER after each integer, and press ENTER twice when done)")

	sli := Input([]int{}, nil)
	BubbleSort(sli)

	fmt.Println("The integers in sorted order:")
	fmt.Print(strings.Trim(fmt.Sprint(sli), "[]"))
}
