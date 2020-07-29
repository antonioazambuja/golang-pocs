package main

import (
	"fmt"

	"github.com/antonioazambuja/golang-pocs/mock_test/calc"
)

func main() {
	fmt.Println("Calculator Golang!")
	fmt.Printf("Sum - 1 + 1: %f\n", calc.NewCalc(1, 1).Sum())
	fmt.Printf("Sub - 1 - 1: %f\n", calc.NewCalc(1, 1).Sub())
	fmt.Printf("Mult - 8 + 2: %f\n", calc.NewCalc(8, 2).Mult())
	fmt.Printf("Div - 8 + 2: %f\n", calc.NewCalc(8, 2).Div())
}
