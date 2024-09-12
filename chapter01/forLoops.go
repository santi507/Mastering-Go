package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i*i, " ")
	}
	fmt.Println()

	// For loop used as while loop
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i*i, " ")
		i++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value", v)
	}
}
