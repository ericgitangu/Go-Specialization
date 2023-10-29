package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var value float64
	radius := 10.0

	/*
	The go routines one that computes area of a circle and the second computes the area. 
	When run concurrently and share a shared resource value, the way value is computed is 
	non deterministic sometimes returning circumference sometimes returning area hence simulating a race condition.
	*/
	go func() {
		value = math.Pi * 2 * radius
		fmt.Println("Circumference: goroutine 1")
	}()

	go func() {
		value = math.Pi * radius * radius
		fmt.Println("Area: goroutine 2")
	}()

	time.Sleep(time.Second)
	fmt.Println(value)
}
