package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		request := strings.Split(input, " ")

		if len(request) != 2 {
			fmt.Println("Please enter a valid animal and a valid verb separated by a space. For example, cow eat, cow move, cow speak, bird eat, bird move, bird speak, snake eat, snake move, snake speak.")
			continue
		}

		switch request[0] {
		case "cow":
			switch request[1] {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println("Verb not recognized for a cow")
			}
		case "bird":
			switch request[1] {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Println("Verb not recognized for a bird")
			}
		case "snake":
			switch request[1] {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println("Verb not recognized for a snake")
			}
		default:
			fmt.Println("Animal not recognized, please enter a valid animal (cow, bird, snake) and a valid verb (eat, move, speak) separated by a space. For example, cow eat, cow move, cow speak, bird eat, bird move, bird speak, snake eat, snake move, snake speak.")
		}
	}
}
