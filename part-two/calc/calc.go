package calc

import (
	"errors"
	"strconv"
)

func haveNextNumber(expression string, idx int) bool {
	if idx+1 != len(expression) {
		switch string(expression[idx+1]) {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".":
			return true
		default:
			return false
		}
	}

	return false
}

func mergeCalculatablesOperations(cArr []iCalculatable, op []string, divisionOrMultiply bool) ([]iCalculatable, []string) {
	for i := 0; i < len(op); i++ {
		clearOp := false
		leftValue := cArr[i]
		rightValue := cArr[i+1]
		tailArray := cArr[i+2:]

		if divisionOrMultiply {
			if op[i] == "*" {
				cArr = append(cArr[:i], &multiply{leftValue, rightValue})
				clearOp = true
			} else if op[i] == "/" {
				cArr = append(cArr[:i], &division{leftValue, rightValue})
				clearOp = true
			}
		} else {
			if op[i] == "+" {
				cArr = append(cArr[:i], &plus{leftValue, rightValue})
				clearOp = true
			} else if op[i] == "-" {
				cArr = append(cArr[:i], &minus{leftValue, rightValue})
				clearOp = true
			}
		}

		if clearOp {
			cArr = append(cArr, tailArray...)
			op = append(op[:i], op[i+1:]...)
			i--
		}
	}

	return cArr, op
}

func parsingExpression(expression string, countOpenBrackets *int, currentIterator int) (iCalculatable, int, error) {
	var calculatableArray []iCalculatable
	var operations []string
	var convertingNumber string
	var curIdx int
	var err error

LOOP:
	for curIdx = currentIterator; curIdx < len(expression); curIdx++ {
		curSymbol := string(expression[curIdx])

		switch curSymbol {

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".":
			if curSymbol == "." && len(convertingNumber) == 0 {
				return nil, -1, errors.New("ошибка в записи выражения")
			}
			convertingNumber += curSymbol
			if !haveNextNumber(expression, curIdx) {
				number, _ := strconv.ParseFloat(convertingNumber, 64)
				calculatableArray = append(calculatableArray, &value{number})
				convertingNumber = ""
			}

		case "+":
			operations = append(operations, "+")

		case "-":
			operations = append(operations, "-")

		case "*":
			operations = append(operations, "*")

		case "/":
			operations = append(operations, "/")

		case "(":
			(*countOpenBrackets)++
			var result iCalculatable
			result, curIdx, err = parsingExpression(expression, countOpenBrackets, curIdx+1)
			if err != nil {
				return nil, -1, err
			}
			calculatableArray = append(calculatableArray, result)

		case ")":
			if *countOpenBrackets == 0 {
				return nil, -1, errors.New("в выражении присутствует лишняя закрывающая скобка")
			}
			(*countOpenBrackets)--
			break LOOP

		}
	}

	if *countOpenBrackets != 0 && curIdx == len(expression) {
		return nil, -1, errors.New("в выражении присуствует лишняя открывающая скобка")
	} else if len(operations)+1 != len(calculatableArray) {
		return nil, -1, errors.New("в выражении присутствуют лишние операции, символы или выражание в скобках пустое")
	}

	calculatableArray, operations = mergeCalculatablesOperations(calculatableArray, operations, true)
	calculatableArray, _ = mergeCalculatablesOperations(calculatableArray, operations, false)

	if len(calculatableArray) != 0 {
		return calculatableArray[0], curIdx, nil
	} else {
		return &value{value: 0}, curIdx, nil
	}
}

func Calc(expression string) (float64, error) {
	countOpenBrackets := 0

	value, _, err := parsingExpression(expression, &countOpenBrackets, 0)

	if err != nil {
		return -1, errors.New("Ошибка парсинга выражения: " + err.Error())
	}

	return value.Calculate(), nil
}
