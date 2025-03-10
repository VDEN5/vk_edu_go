package main

import (
	"fmt"
	"task2/calculator"
)

func main() {
	var expression string
	fmt.Scan(&expression)
	fmt.Printf("%.5f\n", calculator.CalculateExpression(expression))
}
