package main

import (
	"fmt"
	"sync"
)

type ChopSticks struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *ChopSticks
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Philosopher with id: %v starting to eat. \n", p.id)
		fmt.Printf("Philosopher with id: %v finishing eating. \n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

// main function initializes the chopsticks and philosophers, and starts the eating process
func main() {
	// create a slice of 5 chopsticks
	CSticks := make([]*ChopSticks, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopSticks)
	}

	// create a slice of 5 philosophers, each with two chopsticks
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	// create a WaitGroup to wait for all philosophers to finish eating
	var wg sync.WaitGroup
	wg.Add(5)

	// start the eating process for each philosopher in a separate goroutine
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg)
	}

	// wait for all philosophers to finish eating
	wg.Wait()
}
