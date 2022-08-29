package main

import "fmt"

var result int

func main() {
	fmt.Println("--- Inicio ---")

	result = operationMiddleware(sumar)(2, 3)
	fmt.Println(result)
	result = operationMiddleware(restar)(2, 3)
	fmt.Println(result)
	result = operationMiddleware(multiplicar)(2, 3)
	fmt.Println(result)
}

func operationMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}
