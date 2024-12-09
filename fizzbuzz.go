package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args[0] is the name of the command
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: fizzbuzz <number>")
		return
	}

	max, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	for v := range max + 1 {
		if v == 0 {
			continue
		}

		switch {
		case v%15 == 0:
			fmt.Println("FizzBuzz")
		case v%3 == 0:
			fmt.Println("Fizz")
		case v%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(v)
		}
	}
}
