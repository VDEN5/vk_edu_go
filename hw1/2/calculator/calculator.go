package calculator

import (
	"regexp"
	"strconv"
	"strings"
	"task2/stack"
)
/*получаем число со строки+смена позиции, где конец числа*/
func getNumberFromString(s string, pos *int) string {
	var number string
	for ; *pos < len(s); *pos++ {
		_, err := strconv.Atoi(string(s[*pos]))
		if err != nil {
			*pos--
			break
		}
		number += string(s[*pos])
	}
	return number
}
/*очистка выражения, замена на более подхлдящие элементы для парсинга*/
func clean(expression string) string {
	addMult, cleanExpression := regexp.MustCompile("(\\d+)(\\()"), strings.Replace(expression, " ", "", -1)
	cleanExpression = strings.Replace(cleanExpression, ")(", ")*(", -1)
	cleanExpression = addMult.ReplaceAllString(cleanExpression, "${1}*$2")
	return cleanExpression
}
/*определение приоритета исполняемых операций*/
func PriorityCmp(b1, b2 byte) bool {
	priorityMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'_': 3,
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
			res = append(res, getNumberFromString(s, &i))
		case ch == '(':
			st.Push(ch)
		case ch == ')':
			for !st.IsEmpty() && st.Top() != '(' {
				res = append(res, string(st.Pop()))
			}
			st.Pop()
		case ch == '-' && (i == 0 || s[i-1] == '('): //negative number case-swap to _
			ch = '_'
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
/*вычисление польского выражения*/
func CalculateExpression(s string) float64 {
	polishNotation, nums, num1, num2 := getReversePolishNotation(clean(s)), make([]float64, 0, 10), 0.0, 0.0
	if len(polishNotation) <= 2 {
		res, _ := strconv.ParseFloat(polishNotation[0], 64)
		return res
	}
	for _, val := range polishNotation {
		if val == "+" || val == "-" || val == "*" || val == "/" || val == "_" {
			if val == "_" {
				nums[len(nums)-1] = -nums[len(nums)-1]
				continue
			}
			num2, nums = nums[len(nums)-1], nums[:len(nums)-1]
			if len(nums) == 0 {
				num1 = 0
			} else {
				num1 = nums[len(nums)-1]
				nums = nums[:len(nums)-1]
			}
			if val == "+" {
				num1 += num2
			}
			if val == "-" {
				num1 -= num2
			}
			if val == "*" {
				num1 *= num2
			}
			if val == "/" {
				num1 /= num2
			}
			nums = append(nums, num1)
			continue
		}
		num, _ := strconv.ParseFloat(val, 64)
		nums = append(nums, num)
	}
	return nums[0]
}
