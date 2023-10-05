package uniq

import (
	"strconv"
	"strings"
)

func clearingNFields(str string, numberOfFields int) string {
	var changedStr string
	words := strings.Fields(str)

	if numberOfFields >= len(words) {
		return ""
	}

	for i := numberOfFields; i < len(words); i++ {
		changedStr += words[i] + " "
	}

	return changedStr
}

func clearingNChars(str string, numberOfChars int) string {
	if numberOfChars >= len(str) {
		return ""
	} else {
		return str[numberOfChars:]
	}
}

func formatStringWithCurrentFlags(changedStr string, flags map[string]string) string {
	if flags["-f"] != "" {
		numberOfFields, _ := strconv.Atoi(flags["-f"])
		changedStr = clearingNFields(changedStr, numberOfFields)
	}

	if flags["-s"] != "" {
		numberOfChars, _ := strconv.Atoi(flags["-s"])
		changedStr = clearingNChars(changedStr, numberOfChars)
	}

	if flags["-i"] != "" {
		changedStr = strings.ToLower(changedStr)
	}

	return changedStr
}

func comparisonWithFlags(curStr, prevStr string, flags map[string]string, uniqStr *[]string, sequence *bool) {
	// Форматируем текущую и нынешнюю строку для сравнения
	curFormattedStr := formatStringWithCurrentFlags(curStr, flags)
	prevFormattedStr := formatStringWithCurrentFlags(prevStr, flags)

	if flags["-c"] != "" {
		if curFormattedStr == prevFormattedStr {
			sequenceNumber, _ := strconv.Atoi((*uniqStr)[len(*uniqStr)-2])
			(*uniqStr)[len(*uniqStr)-2] = strconv.Itoa(sequenceNumber + 1)
		} else {
			*uniqStr = append(*uniqStr, "1")
			*uniqStr = append(*uniqStr, curStr)
		}
	} else if flags["-d"] != "" {
		// Если мы встретили первые две похожие строки, то другие нет смысла проверять
		// Ставим флаг sequence в положение true
		if curFormattedStr == prevFormattedStr && !(*sequence) {
			*uniqStr = append(*uniqStr, prevStr)
			*sequence = true
		} else if curFormattedStr != prevFormattedStr {
			*sequence = false
		}
	} else if flags["-u"] != "" {
		// Записываем все неравные строки в массив
		// Если встретили похожие, причем предыдушая уже записана в массив, то удаляем ее оттуда
		if curFormattedStr != prevFormattedStr {
			*uniqStr = append(*uniqStr, curStr)
		} else if len(*uniqStr) > 0 {
			uniqFormattedStr := formatStringWithCurrentFlags((*uniqStr)[len(*uniqStr)-1], flags)
			if uniqFormattedStr == curFormattedStr {
				*uniqStr = (*uniqStr)[:len(*uniqStr)-1]
			}
		}
	} else if curFormattedStr != prevFormattedStr {
		// Случай без флагов
		*uniqStr = append(*uniqStr, curStr)
	}
}

func Uniq(preparingStr string, flags map[string]string) ([]string, error) {
	arrayStr := strings.Split(preparingStr, "\n")

	answer := make([]string, 0)
	var curStr string
	var prevStr string
	var sequence bool

	for i := 0; i < len(arrayStr); i++ {
		curStr = arrayStr[i] + "\n"

		comparisonWithFlags(curStr, prevStr, flags, &answer, &sequence)

		prevStr = curStr
	}

	return answer, nil
}
