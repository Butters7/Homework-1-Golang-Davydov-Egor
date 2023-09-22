package uniq

import (
	"fmt"
	"testing"
)

var mainTests = map[string]struct {
	flags  map[string]string
	output []string
}{
	"Test from GitHub lectures: without parametrs": {
		flags:  map[string]string{"inputFile": "test_files/test1.txt"},
		output: []string{"I love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n", "I love music of Kartik.\n"},
	},
	"Test from Github lectures: -c parametr": {
		flags:  map[string]string{"inputFile": "test_files/test1.txt", "-c": "-c"},
		output: []string{"3", "I love music.\n", "1", "\n", "2", "I love music of Kartik.\n", "1", "Thanks.\n", "2", "I love music of Kartik.\n"},
	},
	"Test from Github lectures: -d parametr": {
		flags:  map[string]string{"inputFile": "test_files/test1.txt", "-d": "-d"},
		output: []string{"I love music.\n", "I love music of Kartik.\n", "I love music of Kartik.\n"},
	},
	"Test from GitHub lectures: -u parametr": {
		flags:  map[string]string{"inputFile": "test_files/test1.txt", "-u": "-u"},
		output: []string{"\n", "Thanks.\n"},
	},
	"Test from GitHub lectures: -i parametr": {
		flags:  map[string]string{"inputFile": "test_files/test2.txt", "-i": "-i"},
		output: []string{"I LOVE MUSIC.\n", "\n", "I love MuSIC of Kartik.\n", "Thanks.\n", "I love music of kartik.\n"},
	},
	"Test from GitHub lectures: -f num parametr": {
		flags:  map[string]string{"inputFile": "test_files/test3.txt", "-f": "1"},
		output: []string{"We love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n"},
	},
	"Test from GitHub lectures: -s num parametr": {
		flags:  map[string]string{"inputFile": "test_files/test4.txt", "-s": "1"},
		output: []string{"I love music.\n", "\n", "I love music of Kartik.\n", "We love music of Kartik.\n", "Thanks.\n"},
	},
	"Test with 3 parametres #1": {
		flags:  map[string]string{"inputFile": "test_files/test5.txt", "-u": "-u", "-i": "-i", "-f": "1", "-s": "1"},
		output: []string{"\n", "The tea is hot!\n", "This tea is not hot.\n"},
	},
	"Test with 3 parametres #2": {
		flags:  map[string]string{"inputFile": "test_files/test5.txt", "-d": "-d", "-i": "-i", "-f": "1", "-s": "1"},
		output: []string{"I dove music.\n", "Thanks.\n"},
	},
	"Test with 3 parametres #3": {
		flags:  map[string]string{"inputFile": "test_files/test6.txt", "-c": "-c", "-i": "-i", "-f": "2", "-s": "2"},
		output: []string{"2", "I have a friend\n", "4", "\n", "3", "I love Golang\n"},
	},
	"Test with empty file": {
		flags:  map[string]string{"inputFile": "test_files/test7.txt"},
		output: []string{},
	},
}

func TestMain(t *testing.T) {
	for name, test := range mainTests {

		result, err := comparison(test.flags)
		if err != nil || len(result) != len(test.output) {
			t.Fatalf("Test (%s) failed!\n", name)
		}

		for i := 0; i < len(result); i++ {
			if result[i] != test.output[i] {
				t.Fatalf("Test (%s) failed", name)
			}
		}

		fmt.Printf("Test %s:\t OK\n", name)
	}
}
