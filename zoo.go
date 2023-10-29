package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

type Bird struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

type Snake struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Split(input, " ")

		if len(args) != 3 {
			fmt.Println("Invalid command sequence, expected three arguments.")
			continue
		}

		command := args[0]
		name := args[1]
		typeOfAnimal := args[2]

		switch command {
		case "newanimal":
			switch typeOfAnimal {
			case "cow":
				animals[name] = Cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = Bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = Snake{name, "mice", "slither", "hsss"}
			default:
				fmt.Println("Invalid animal type. Please enter a valid animal type.")
			}
			fmt.Println("Created it!")
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Println("Animal not found. Try creating it first with the command 'newanimal'.")
				continue
			}

			switch typeOfAnimal {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid animal verb requested. Please enter a valid animal verb.")
			}
		default:
			fmt.Println("Invalid command. Valid commands are 'newanimal' and 'query.")
		}
	}
}
