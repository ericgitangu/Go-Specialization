package main

import (
	"fmt"
	"sort"
)

func main() {
	// take input of 3 distinct integers
	var num1, num2, num3, num4 int
	fmt.Println("Enter 3 distinct integers:")
	fmt.Scan(&num1, &num2, &num3)

	// create a slice with the input integers
	slice := []int{num1, num2, num3}

	// sort the slice
	sort.Ints(slice)

	// print the sorted slice
	fmt.Println("Sorted slice:", slice)

	// take input of 4 distinct integers
	fmt.Println("Enter 4 distinct integers:")
	fmt.Scan(&num1, &num2, &num3, &num4)

	// create a slice from the input
	slice = []int{num1, num2, num3, num4}

	// sort the slice
	sort.Ints(slice)

	// print the sorted slice
	fmt.Println("Sorted slice:", slice)
}
