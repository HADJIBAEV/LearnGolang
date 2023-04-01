package main

import "fmt"

func main() {
	// _____ If else example _____
	var number = [...]int{1, 2, 3, 4}
	numbers := [...]int{1, 2, 3, 4}
	if number == numbers {
		fmt.Println(number, " == ", numbers)
	} else {
		fmt.Println(number)
		fmt.Println(numbers)
	}

	a := 12
	if a > 0 {
		fmt.Println(a, "> 0")
	} else if a != 12 {
		fmt.Println("A != 12")
	} else {
		fmt.Println("A = 12")
	}

	fmt.Println("Factorial 4 = ", Factorial(4))
	fmt.Println("Fibonachi = ", Fibonacci(6))

	// _____ Switch example _____
	b := 12
	switch b {
	case 9, 10, 11, 12:
		fmt.Println("SwitchCase B = 9")
	case 8:
		fmt.Println("SwitchCase B = 8")
	case 7:
		fmt.Println("SwitchCase B = 7")
	default:
		fmt.Println("Default variable SwitchCase")
	}

	// _____  For Loop example:

	// _____  For Loop with Condition example:

	// _____  Continue & Break example:

	// _____  Struct example:

	// _____  Class example:

}

// Factorial _____________
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci _____________
func Fibonacci(n uint) []uint {
	var sl = make([]uint, n)
	var i uint = 1
	for i < n {
		sl[i] = fibonacciOne(i)
		i++
	}
	return sl
}
func fibonacciOne(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacciOne(n-1) + fibonacciOne(n-2)
	}
}

// _____________
