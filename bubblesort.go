package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 10)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 10 integers:")
	for i := 0; i < 10; i++ {
		scanner.Scan()
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			i--
			continue
		}
		nums[i] = num
	}
	fmt.Println("Before sorting:", nums)
	BubbleSort(nums)
	fmt.Println("After sorting:", nums)
}

// BubbleSort sorts a slice of integers using the bubble sort algorithm
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
			}
		}
	}
}

// swap swaps two elements in a slice of integers
func swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
