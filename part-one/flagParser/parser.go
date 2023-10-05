package flagparser

import (
	"errors"
	"flag"
	"strconv"
)

func ParsingCommandArguments() (map[string]string, error) {
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

	lengthRemainingFlags := len(flag.Args())
	if lengthRemainingFlags > 2 {
		return flags, errors.New("слишком много аргументов")
	} else if lengthRemainingFlags == 1 {
		flags["inputFile"] = flag.Arg(0)
	} else if lengthRemainingFlags == 2 {
		flags["inputFile"] = flag.Arg(0)
		flags["outputFile"] = flag.Arg(1)
	}

	return flags, nil
}
