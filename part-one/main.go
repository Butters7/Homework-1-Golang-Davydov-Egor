package main

import (
	flagparser "HW1/part-one/flagParser"
	"HW1/part-one/uniq"
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Считываем данные из текста или STDIN
func getStrFromBuffer(filePath string) (string, error) {
	var resultString string

	var reader *bufio.Reader

	if filePath != "" {
		file, err := os.Open(filePath)

		if err != nil {
			return "", errors.New("Error: " + err.Error())
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

		resultString += curStr
	}

	return resultString, nil
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

// Вывод в соответствии с флагом outputFile
func writeResult(result []string, flags map[string]string) {
	if flags["outputFile"] == "" {
		consoleWriter(result, flags)
	} else {
		fileWriter(result, flags)
	}
}

func main() {
	flags, err := flagparser.ParsingCommandArguments()
	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}

	preparingStr, err := getStrFromBuffer(flags["inputFile"])
	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}

	answer, err := uniq.Uniq(preparingStr, flags)
	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}

	writeResult(answer, flags)
}
