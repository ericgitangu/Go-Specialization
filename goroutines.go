package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Prompt user to input a series of integers
	var input string
	fmt.Println("Enter a series of integers separated by a space:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	// Convert input string to slice of integers
	var nums []int
	for _, n := range strings.Fields(input) {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		nums = append(nums, num)
	}

	// Partition the array into 4 parts
	partSize := len(nums) / 4
	parts := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			end = len(nums)
		}
		parts[i] = nums[start:end]
	}

	// Print the subarray that each goroutine will sort
	for i, part := range parts {
		wg.Add(1)
		go func(i int, part []int) {
			defer wg.Done() // Defer statement to signal the WaitGroup that the goroutine has finished
			fmt.Printf("Sorting Part %d: %v\n", i+1, part)
			sort.Ints(part)
		}(i, part)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Merge the 4 sorted subarrays into one large sorted array
	var merged []int
	for _, part := range parts {
		merged = merge(merged, part)
	}

	// Print the entire sorted list
	fmt.Println("Sorted List:", merged)
}

// Merge two sorted arrays into one sorted array
func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		merged[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		merged[k] = b[j]
		j++
		k++
	}
	return merged
}
