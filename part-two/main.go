package main

import (
	"HW1/part-two/calc"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		fmt.Println("Неверное количество аргументов")
		return
	}

	expression := args[1]
	result, err := calc.Calc(expression)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
