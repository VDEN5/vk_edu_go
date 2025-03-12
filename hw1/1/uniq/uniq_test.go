package uniq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testUniq struct {
	input   []string
	options Options
	output  []string
}

func TestCalculateExpression(t *testing.T) {
	testData := []testUniq{{
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		Options{},
		[]string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
		{
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			Options{
				c: true,
			},
			[]string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik.",
			},
		},
		{
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			Options{
				d: true,
			},
			[]string{
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		{
			[]string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			Options{
				u: true,
			},
			[]string{
				"",
				"Thanks.",
			},
		},
		{
			[]string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
			Options{
				i: true,
			},
			[]string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
				"I love music of kartik.",
			},
		},
		{
			[]string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			Options{
				f: 1,
			},
			[]string{
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			[]string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			Options{
				ch: 1,
			},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		}}
	for _, test := range testData {
		res, _ := Uniq(test.input, test.options)
		require.Equal(t, test.output, res)
	}
}

var uniqTestsNegative = []struct {
	input   []string
	options Options
	output  []string
}{
	{
		[]string{"bjkrnl njylmkg;f, bjhk.,,,,,,,,,"},
		Options{
			c: true,
			d: true,
		},
		[]string{},
	},
}

func TestNegative(t *testing.T) {
	for _, tt := range uniqTestsNegative {
		result, err := Uniq(tt.input, tt.options)
		fmt.Println(err, tt.options)
		require.Error(t, err)
		require.Equal(t, tt.output, result)
	}
}
