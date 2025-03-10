package calculator

import (
	"testing"
)

type testCalc struct {
	expression string
	answer     float64
}

func TestCalculateExpression(t *testing.T) {
	testData := []testCalc{
		//base tests
		{
			expression: "(1+2)-3",
			answer:     0,
		},
		{
			expression: "(1+2)*3",
			answer:     9,
		},
		//my tests
		{
			expression: "(1+2)/3",
			answer:     1,
		},
		{
			expression: "3*2/3",
			answer:     2,
		},
		{
			expression: "-10*(-5*5)",
			answer:     250,
		},
		//teacher`s test
		{
			expression: "-(-11-(1*20/2)-11/2*3)",
			answer:     37.5,
		},
	}
	for _, test := range testData {
		answer := CalculateExpression(test.expression)
		if answer != test.answer {
			t.Errorf("calculateExpression(%s) => %f, want %f", test.expression, answer, test.answer)
		}
	}
}
