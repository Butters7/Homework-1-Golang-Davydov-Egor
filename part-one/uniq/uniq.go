package uniq

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsingCommandArguments() (map[string]string, error) {
	flags := make(map[string]string)

	var cFl, dFl, uFl, iFl bool

	flag.BoolVar(&cFl, "c", false, "Counting line repetition")
	flag.BoolVar(&dFl, "d", false, "Output of duplicate lines")
	flag.BoolVar(&uFl, "u", false, "Output of non-repeating lines")
	flag.BoolVar(&iFl, "i", false, "Ignore case")

	var fFl, sFl int

	flag.IntVar(&fFl, "f", 0, "Ignore N words")
	flag.IntVar(&sFl, "s", 0, "Ignore N symbols")

	flag.Parse()

	counter := 0
	if cFl {
		counter++
		flags["-c"] = "-c"
	}
	if dFl {
		counter++
		flags["-d"] = "-d"
	}
	if uFl {
		counter++
		flags["-u"] = "-u"
	}

	if counter > 1 {
		return flags, errors.New("флаги -c -d -u взаимозаменяемые и не используются вместе")
	}

	if iFl {
		flags["-i"] = "-i"
	}

	if fFl != 0 {
		flags["-f"] = strconv.Itoa(fFl)
	}

	if sFl != 0 {
		flags["-s"] = strconv.Itoa(sFl)
	}

	remainingFlags := flag.Args()
	getInputFile := false
	getOutputFile := false
	for i := 0; i < len(remainingFlags); i++ {
		if string(remainingFlags[i][0]) == "-" {
			return flags, errors.New("неопознанный флаг: " + remainingFlags[i])
		} else if !getInputFile {
			flags["inputFile"] = remainingFlags[i]
			getInputFile = true
		} else if !getOutputFile {
			flags["outputFile"] = remainingFlags[i]
			getOutputFile = true
		} else {
			return flags, errors.New("слишком много аргументов")
		}
	}

	return flags, nil
}

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
		if curFormattedStr == prevFormattedStr && !(*sequence) {
			*uniqStr = append(*uniqStr, prevStr)
			*sequence = true
		} else if curFormattedStr != prevFormattedStr {
			*sequence = false
		}
	} else if flags["-u"] != "" {
		if curFormattedStr != prevFormattedStr {
			*uniqStr = append(*uniqStr, curStr)
		} else if len(*uniqStr) > 0 {
			uniqFormattedStr := formatStringWithCurrentFlags((*uniqStr)[len(*uniqStr)-1], flags)
			if uniqFormattedStr == curFormattedStr {
				*uniqStr = (*uniqStr)[:len(*uniqStr)-1]
			}
		}
	} else if curFormattedStr != prevFormattedStr {
		*uniqStr = append(*uniqStr, curStr)
	}
}

func fileWriter(content []string, flags map[string]string) {
	file, err := os.Create(flags["outputFile"])

	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Ошибка: ", err.Error())
		}
	}()

	writer := bufio.NewWriter(file)

	for i := 0; i < len(content); i++ {
		if flags["-c"] != "" {
			writer.WriteString(content[i] + " " + content[i+1])
			i++
		} else {
			writer.WriteString(content[i])
		}
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Ошибка: ", err.Error())
	}
}

func consoleWriter(content []string, flags map[string]string) {
	for i := 0; i < len(content); i++ {
		if flags["-c"] != "" {
			fmt.Print(content[i] + " " + content[i+1])
			i++
		} else {
			fmt.Print(content[i])
		}
	}
}

func writeResult(result []string, flags map[string]string) {
	if flags["outputFile"] == "" {
		consoleWriter(result, flags)
	} else {
		fileWriter(result, flags)
	}
}

func comparison(flags map[string]string) ([]string, error) {
	var reader *bufio.Reader

	answer := make([]string, 0)
	var prevStr string
	var sequence bool

	if flags["inputFile"] != "" {
		file, err := os.Open(flags["inputFile"])

		if err != nil {
			return answer, errors.New("Error: " + err.Error())
		}

		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("Error", err.Error())
			}
		}()

		reader = bufio.NewReader(file)

	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	for {
		curStr, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		comparisonWithFlags(curStr, prevStr, flags, &answer, &sequence)

		prevStr = curStr
	}

	return answer, nil
}

func Uniq(args []string) {
	var answer []string

	flags, err := parsingCommandArguments()
	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}

	answer, err = comparison(flags)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	writeResult(answer, flags)
}
