package main

import (
	"fmt"
)

func main() {
	fib := []int{1, 1, 2, 3, 5, 8}
	Demo(fib...)
	fmt.Println(Sum(fib...))

}

func Sum(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
