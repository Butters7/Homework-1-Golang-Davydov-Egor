package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert := assert.New(t)
	for name, test := range mainTests {

		result, err := Calc(test.input)

		assert.Nil(err, "Test (%s) failed.\n", name)
		assert.Equal(result, test.output, "Test %s failed. Expected: %f. Current: %f\n", name, test.output, result)
	}
}
