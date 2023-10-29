package main

import "fmt"

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Print("Enter acceleration in m**2: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity in m/s: ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement in m: ")
	fmt.Scan(&so)

	fmt.Print("Enter time in s: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Printf("Displacement after %g seconds is %g meters.\n", t, fn(t))
}
