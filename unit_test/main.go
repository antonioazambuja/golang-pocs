package main

import (
	"fmt"

	"github.com/antonioazambuja/golang-pocs/unit_test/calc"
)

func main() {
	fmt.Println("Calculator Golang!")
	fmt.Printf("Sum - 1 + 1: %f\n", calc.Sum(1, 1))
	fmt.Printf("Sub - 1 - 1: %f\n", calc.Sub(1, 1))
	fmt.Printf("Mult - 8 + 2: %f\n", calc.Mult(8, 2))
	fmt.Printf("Div - 8 + 2: %f\n", calc.Div(8, 2))
}
