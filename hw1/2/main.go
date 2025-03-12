package main

import (
	"fmt"
	"task2/calculator"
)

func main() {
	var expression string
	fmt.Scan(&expression)
	res, _ := calculator.CalculateExpression(expression)
	fmt.Printf("%.5f\n", res)
}
