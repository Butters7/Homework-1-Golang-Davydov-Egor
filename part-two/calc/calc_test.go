package calc

import (
	"fmt"
	"testing"
)

var mainTests = map[string]struct {
	input  string
	output float64
}{
	"Testing \"5+2\"": {
		input:  "5+2",
		output: 7,
	},
	"Testing \"5-2\"": {
		input:  "5-2",
		output: 3,
	},
	"Testing \"5*2\"": {
		input:  "5*2",
		output: 10,
	},
	"Testing \"5/2\"": {
		input:  "5/2",
		output: 2.5,
	},
	"Testing \"(1+2)-3\"": {
		input:  "(1+2)-3",
		output: 0,
	},
	"Testing \"(1+2)*3\"": {
		input:  "(1+2)*3",
		output: 9,
	},
	"Testing \"5+2*(3+4*2+1/2)-(5*2)+1\"": {
		input:  "5+2*(3+4*2+1/2)-(5*2)+1",
		output: 19,
	},
	"Testing \"25/5.0\"": {
		input:  "25/5.0",
		output: 5,
	},
	"Testing \"(((((124 + (5.2 * 5 * (0 + 1 * 1)))))))\"": {
		input:  "(((((124 + ((5.2 * 5 * ((0 + 1 * 1)))))))))",
		output: 150,
	},
}

func TestMain(t *testing.T) {
	for name, test := range mainTests {
		result, err := Calc(test.input)

		if err != nil {
			t.Fatalf("Test %s failed. %s\n", name, err.Error())
		} else if result != test.output {
			t.Fatalf("Test %s failed. Result are not equal. Expected: %f. Current: %f\n", name, test.output, result)
		}

		fmt.Printf("Test %s passed\n", name)
	}
}
