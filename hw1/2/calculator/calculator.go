package calculator

import (
	"regexp"
	"strconv"
	"strings"
	"task2/stack"
)

/*получаем число со строки+смена позиции, где конец числа*/
func getNumberFromString(s string, pos int) (string, int) {
	var number string
	for ; pos < len(s); pos++ {
		_, err := strconv.Atoi(string(s[pos]))
		if err != nil {
			pos--
			break
		}
		number += string(s[pos])
	}
	return number, pos
}

/*очистка выражения, замена на более подхлдящие элементы для парсинга*/
func clean(expression string) string {
	addMult, cleanExpression := regexp.MustCompile("(\\d+)(\\()"), strings.Replace(expression, " ", "", -1)
	cleanExpression = strings.Replace(cleanExpression, ")(", ")*(", -1)
	cleanExpression = addMult.ReplaceAllString(cleanExpression, "${1}*$2")
	return cleanExpression
}

const (
	PLUS  = '+'
	MINUS = '-'
	MUL   = '*'
	DIV   = '/'
	NEG   = '_'
)

/*определение приоритета исполняемых операций*/
func PriorityCmp(b1, b2 byte) bool {
	priorityMap := map[byte]int{
		PLUS:  1,
		MINUS: 1,
		MUL:   2,
		DIV:   2,
		NEG:   3,
	}
	return priorityMap[b1] >= priorityMap[b2]
}

/*алгоритм преобразования в польскую запись, посредством Дейкстры*/
func getReversePolishNotation(s string) []string {
	st, res := new(stack.Stack), make([]string, 0)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		_, err := strconv.Atoi(string(ch))
		switch {
		case err == nil:
			num, newPos := getNumberFromString(s, i)
			res = append(res, num)
			i = newPos
		case ch == '(':
			st.Push(ch)
		case ch == ')':
			for !st.IsEmpty() && st.Top() != '(' {
				res = append(res, string(st.Pop()))
			}
			st.Pop()
		case ch == MINUS && (i == 0 || s[i-1] == '('): //negative number case-swap to _
			ch = NEG
			for !st.IsEmpty() && PriorityCmp(st.Top(), ch) {
				res = append(res, string(st.Pop()))
			}
			st.Push(ch)
		default:
			for !st.IsEmpty() && PriorityCmp(st.Top(), ch) {
				res = append(res, string(st.Pop()))
			}
			st.Push(ch)
		}
	}
	for !st.IsEmpty() {
		res = append(res, string(st.Pop()))
	}
	return res
}

/*ВАЖНО!!! использую, когда гарантированно, что можно распарсить*/
func StrToFloat64(str string) float64 {
	res, _ := strconv.ParseFloat(str, 64)
	return res
}

/*вычисление польского выражения*/
func CalculateExpression(s string) float64 {
	polishNotation, numStack, topStack1, topStack2 := getReversePolishNotation(clean(s)), make([]float64, 0, 10), 0.0, 0.0
	if len(polishNotation) <= 2 {
		return StrToFloat64(polishNotation[0])
	}
	for _, val := range polishNotation {
		if v := val[0]; v == PLUS || v == MINUS || v == MUL || v == DIV || v == NEG {
			if v == NEG {
				numStack[len(numStack)-1] = -numStack[len(numStack)-1]
				continue
			}
			topStack1, topStack2, numStack = 0, numStack[len(numStack)-1], numStack[:len(numStack)-1]
			if len(numStack) > 0 {
				topStack1 = numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
			}
			switch v {
			case PLUS:
				topStack1 += topStack2
			case MINUS:
				topStack1 -= topStack2
			case MUL:
				topStack1 *= topStack2
			case DIV:
				topStack1 /= topStack2
			}
			numStack = append(numStack, topStack1)
			continue
		}
		numStack = append(numStack, StrToFloat64(val))
	}
	return numStack[0]
}
