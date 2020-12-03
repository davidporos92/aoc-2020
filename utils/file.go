package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func MustReadStringMapFromFile(path string, sep byte, trim string) (myMap [][]string) {
	MustReadFile(path, sep, trim, func(line string) {
		lineSlice := make([]string, len(line))
		for i, char := range []rune(line) {
			lineSlice[i] = string(char)
		}
		myMap = append(myMap, lineSlice)
	})

	return myMap
}

func MustReadIntSliceFromFile(path string, sep byte, trim string) []int {
	lines := make([]int, 0)

	MustReadFile(path, sep, trim, func(line string) {
		entry, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("String to Int conversion error: %v", err))
		}

		lines = append(lines, entry)
	})

	return lines
}

func MustReadFile(path string, sep byte, trim string, cb func(line string)) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot open file: %v", err))
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString(sep)
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("Cannot read line: %v", err))
		}

		line = strings.Trim(line, trim)

		if line != "" {
			cb(line)
		}

		if err == io.EOF {
			break
		}
	}
}
