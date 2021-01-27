package main

import (
	"fmt"
	"sort"
	"sync"
)

func Input(sli []int, err error) []int {
	if err != nil {
		return sli
	}
	var new int
	n, err := fmt.Scanf("%d", &new)
	if n == 1 {
		sli = append(sli, new)
	}
	return Input(sli, err)
}

func Sort(wg *sync.WaitGroup, sli []int, ch chan []int) {
	sort.Ints(sli)
	fmt.Println(sli)
	ch <- sli
	wg.Done()
}

func main() {

	fmt.Println("Please, type in a sequence of integers: ")
	fmt.Println("(press ENTER after each integer, and press ENTER twice when done)")
	input := Input([]int{}, nil)

	var wg sync.WaitGroup
	ch := make(chan []int, 4)
	wg.Add(4)
	quarter := int(len(input)/4)

	go Sort(&wg, input[:quarter], ch)
	go Sort(&wg, input[quarter:quarter*2], ch)
	go Sort(&wg, input[quarter*2:quarter*3], ch)
	go Sort(&wg, input[quarter*3:], ch)

	wg.Wait()
	close(ch)
	var result []int
	for sli := range ch {
		result = append(result, sli...)
	}
	sort.Ints(result)
	fmt.Println(result)
}
