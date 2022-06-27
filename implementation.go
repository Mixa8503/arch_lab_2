package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack []float64

func (s *Stack) push(item float64) {
	*s = append(*s, item)
}

func (s *Stack) pop() (float64, error) {
	l := len(*s)
	if l == 0 {
		return 0.0, fmt.Errorf("Stack is empty")
	}
	item := (*s)[l-1]
	*s = (*s)[:l-1]
	return item, nil
}

func isOperator(str string) bool {
	operators := [...]string{"+", "-", "*", "/", "^"}
	for _, operator := range operators {
		if str == operator {
			return true
		}
	}
	return false
}

func Prefix_Val(expr []string) bool {
	if !isOperator(expr[0]) {
		return false
	}
	var Operator_counter, Operand_counter int
	for _, str := range expr {
		if isOperator(str) {
			Operator_counter++
			continue
		}
		Operand_counter++
	}
	if Operand_counter == Operator_counter+1 {
		return true
	}
	return false
}

func Prefix_Eval(expr string) (float64, error) {
	var stack Stack
	values := strings.Fields(expr)
	res := Prefix_Val(values)
	if !res {
		return 0.0, fmt.Errorf("incorrect input")
	}
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		if num, err := strconv.ParseFloat(value, 64); err == nil {
			stack.push(num)
		} else {
			operand1, err1 := stack.pop()
			operand2, err2 := stack.pop()
			if err1 != nil || err2 != nil {
				return 0.0, fmt.Errorf("Something is wrong, you need to check input")
			}
			switch value {
			case "+":
				stack.push(operand1 + operand2)
			case "-":
				stack.push(operand1 - operand2)
			case "*":
				stack.push(operand1 * operand2)
			case "/":
				stack.push(operand1 / operand2)
			case "^":
				stack.push(math.Pow(operand1, operand2))
			default:
				return 0.0, fmt.Errorf("Choose operators: { +, -, *, /, ^ }")
			}
		}
	}
	return stack.pop()
}
