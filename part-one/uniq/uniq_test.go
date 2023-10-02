package uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainTests = map[string]struct {
	preparingStr string
	flags        map[string]string
	output       []string
}{
	"Test from GitHub lectures: without parametrs": {
		preparingStr: "I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
		flags:        map[string]string{},
		output:       []string{"I love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n", "I love music of Kartik.\n"},
	},
	"Test from Github lectures: -c parametr": {
		preparingStr: "I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
		flags:        map[string]string{"-c": "-c"},
		output:       []string{"3", "I love music.\n", "1", "\n", "2", "I love music of Kartik.\n", "1", "Thanks.\n", "2", "I love music of Kartik.\n"},
	},
	"Test from Github lectures: -d parametr": {
		preparingStr: "I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
		flags:        map[string]string{"-d": "-d"},
		output:       []string{"I love music.\n", "I love music of Kartik.\n", "I love music of Kartik.\n"},
	},
	"Test from GitHub lectures: -u parametr": {
		preparingStr: "I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
		flags:        map[string]string{"-u": "-u"},
		output:       []string{"\n", "Thanks.\n"},
	},
	"Test from GitHub lectures: -i parametr": {
		preparingStr: "I LOVE MUSIC.\nI love music.\nI LoVe MuSiC.\n\nI love MuSIC of Kartik.\nI love music of kartik.\nThanks.\nI love music of kartik.\nI love MuSIC of Kartik.",
		flags:        map[string]string{"-i": "-i"},
		output:       []string{"I LOVE MUSIC.\n", "\n", "I love MuSIC of Kartik.\n", "Thanks.\n", "I love music of kartik.\n"},
	},
	"Test from GitHub lectures: -f num parametr": {
		preparingStr: "We love music.\nI love music.\nThey love music.\n\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.",
		flags:        map[string]string{"-f": "1"},
		output:       []string{"We love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n"},
	},
	"Test from GitHub lectures: -s num parametr": {
		preparingStr: "I love music.\nA love music.\nC love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.",
		flags:        map[string]string{"-s": "1"},
		output:       []string{"I love music.\n", "\n", "I love music of Kartik.\n", "We love music of Kartik.\n", "Thanks.\n"},
	},
	"Test with 3 parametres #1": {
		preparingStr: "I dove music.\nWe love music.\nThey nove music.\n\nThe tea is hot!\nThis tea is not hot.\nThanks.\nThanks.",
		flags:        map[string]string{"-u": "-u", "-i": "-i", "-f": "1", "-s": "1"},
		output:       []string{"\n", "The tea is hot!\n", "This tea is not hot.\n"},
	},
	"Test with 3 parametres #2": {
		preparingStr: "I dove music.\nWe love music.\nThey nove music.\n\nThe tea is hot!\nThis tea is not hot.\nThanks.\nThanks.",
		flags:        map[string]string{"-d": "-d", "-i": "-i", "-f": "1", "-s": "1"},
		output:       []string{"I dove music.\n", "Thanks.\n"},
	},
	"Test with 3 parametres #3": {
		preparingStr: "I have a friend\nAlice has a friend\n\nCockroach\nZamkroach\n\nI love Golang\nI LoVE gOLANg\nWe love GolANG",
		flags:        map[string]string{"-c": "-c", "-i": "-i", "-f": "2", "-s": "2"},
		output:       []string{"2", "I have a friend\n", "4", "\n", "3", "I love Golang\n"},
	},
	"Test with empty string": {
		preparingStr: "",
		flags:        map[string]string{},
		output:       []string{"\n"},
	},
}

func TestUniq(t *testing.T) {
	assert := assert.New(t)
	for name, test := range mainTests {

		result, err := Uniq(test.preparingStr, test.flags)

		assert.Nil(err, "Test (%s) failed!\n", name)
		assert.Equal(len(result), len(test.output), "Test (%s) failed!\n", name)

		for i := 0; i < len(result); i++ {
			assert.Equal(result[i], test.output[i], "Test (%s) failed!\n", name)
		}
	}
}
