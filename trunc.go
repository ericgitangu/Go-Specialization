package main

import (
	"fmt"
)

func main() {
	var num float64

	for i := 0; i < 2; i++ {
		_, error := fmt.Scan(&num)
		if(error == nil) {
			fmt.Println(int(num))
		}
	}
}
