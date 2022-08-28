package main

import "fmt"

func main() {
	prices := []int{98, 125, 232, 147, 486}

	for i := 0; i < len(prices); i++ {

		fmt.Printf("%d円\n", prices[i])
	}
}
