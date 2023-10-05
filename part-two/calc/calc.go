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

// Итерируемся по массиву. Натыкаясь на мат. операцию, смотрим, какие значения в массиве iCalucaltable. Берем левое и правое
// Затем они удаляются из массива, по итогу получая вместо 2 элементов iCalculatable получается один
func mergeCalculatablesOperations(cArr []iCalculatable, op []string, divisionOrMultiply bool) ([]iCalculatable, []string, error) {
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
				if rightValue.Calculate() == 0 {
					return nil, nil, errors.New("деление на ноль")
				}
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

	return cArr, op, nil
}

func parsingExpression(expression string, countOpenBrackets int, currentIterator int) (iCalculatable, int, error) {
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
			// При парсинге цифры смотрим следующий символ.
			// Если он является точкой или цифрой, то парсим дальше, иначе добавляем в массив
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
			// Следующее выражение в скобках обрабатывается как обычное выражение без скобок с созданием нового массива и конечным значение
			countOpenBrackets++
			var result iCalculatable
			result, curIdx, err = parsingExpression(expression, countOpenBrackets, curIdx+1)
			if err != nil {
				return nil, -1, err
			}
			countOpenBrackets--
			calculatableArray = append(calculatableArray, result)

		case ")":
			if countOpenBrackets == 0 {
				return nil, -1, errors.New("в выражении присутствует лишняя закрывающая скобка")
			}
			break LOOP

		}
	}

	// Так как все опирации бинарные, то общее кол-во операций должно быть на один меньше всех цифр
	if countOpenBrackets != 0 && curIdx == len(expression) {
		return nil, -1, errors.New("в выражении присуствует лишняя открывающая скобка")
	} else if len(operations)+1 != len(calculatableArray) {
		return nil, -1, errors.New("в выражении присутствуют лишние операции, символы или выражание в скобках пустое")
	}

	// Получив операции и iСalculatable выражение сливаем все в одно значение.
	// Первый приоритет для умножения и деления
	calculatableArray, operations, err = mergeCalculatablesOperations(calculatableArray, operations, true)

	if err != nil {
		return nil, -1, err
	}

	// Второй для сложения и вычитания
	calculatableArray, _, _ = mergeCalculatablesOperations(calculatableArray, operations, false)

	// Если в конце слияния получилось так, что массив пустой, то значит, что пользователь написа пустое выражение в скобках
	// () = 0
	if len(calculatableArray) != 0 {
		return calculatableArray[0], curIdx, nil
	} else {
		return &value{value: 0}, curIdx, nil
	}
}

func Calc(expression string) (float64, error) {
	//Держим "в уме" количество открытых скобок, чтобы слить их с закрытыми
	countOpenBrackets := 0

	value, _, err := parsingExpression(expression, countOpenBrackets, 0)

	if err != nil {
		return -1, errors.New("Ошибка парсинга выражения: " + err.Error())
	}

	return value.Calculate(), nil
}
