package main

import "fmt"

func main() {
	intarray := [5]int{98, 125, 232, 147, 486}
	slice13 := intarray[1:4] // 1番目から3番目まで
	slice0 := intarray[:1]   // 0番目
	slice4 := intarray[4:]   // 4番目から、4番目まで
	slice13[2] -= 100

	fmt.Printf("slice13[1]は%d\n", slice13[1])
	fmt.Printf("slice0の要素数は%d\n", len(slice0))
	fmt.Printf("slice4[0]の値は%d\n", slice4[0])
	fmt.Printf("今やintarray[3]の値は%d\n", intarray[3])
}
