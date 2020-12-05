package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type batchSeparator uint

const (
	BatchSeparatorBlankLine batchSeparator = iota
)

type Reader interface {
	SetValueSeparator(byte) Reader
	SetTrimCutSet(string) Reader
	SetIgnoreBlankLines(bool) Reader

	MustReadStringBatchesFromFile(batchSeparator batchSeparator) []string
	MustReadStringMapFromFile() [][]string
	MustReadIntSliceFromFile() []int
	MustReadFile(func(string))
}

type reader struct {
	path             string
	valueSeparator   byte
	trimCutSet       string
	ignoreBlankLines bool
}

func NewReader(path string) Reader {
	return &reader{
		path:             path,
		valueSeparator:   '\n',
		trimCutSet:       "\n\r",
		ignoreBlankLines: true,
	}
}

func (r *reader) SetValueSeparator(valueSeparator byte) Reader {
	r.valueSeparator = valueSeparator

	return r
}

func (r *reader) SetTrimCutSet(trimCutSet string) Reader {
	r.trimCutSet = trimCutSet

	return r
}

func (r *reader) SetIgnoreBlankLines(ignore bool) Reader {
	r.ignoreBlankLines = ignore

	return r
}

func (r *reader) MustReadStringBatchesFromFile(batchSeparator batchSeparator) []string {
	if batchSeparator == BatchSeparatorBlankLine {
		r.SetIgnoreBlankLines(false)
	}

	batch := make([]string, 0)
	var currentBatch []string

	r.MustReadFile(func(line string) {
		switch batchSeparator {
		case BatchSeparatorBlankLine:
			if line == "" {

				batch = append(batch, strings.Join(currentBatch, " "))
				currentBatch = make([]string, 0)
				return
			}
		}

		currentBatch = append(currentBatch, line)
	})

	return append(batch, strings.Join(currentBatch, " "))
}

func (r *reader) MustReadStringMapFromFile() (myMap [][]string) {
	r.MustReadFile(func(line string) {
		lineSlice := make([]string, len(line))
		for i, char := range []rune(line) {
			lineSlice[i] = string(char)
		}
		myMap = append(myMap, lineSlice)
	})

	return myMap
}

func (r *reader) MustReadIntSliceFromFile() []int {
	lines := make([]int, 0)

	r.MustReadFile(func(line string) {
		entry, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("String to Int conversion error: %v", err))
		}

		lines = append(lines, entry)
	})

	return lines
}

func (r *reader) MustReadFile(cb func(line string)) {
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("Cannot open file: %v", err))
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString(r.valueSeparator)
		if err != nil && err != io.EOF {
			panic(fmt.Sprintf("Cannot read line: %v", err))
		}
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, r.trimCutSet)

		if r.ignoreBlankLines && line == "" {
			continue
		}

		cb(line)
	}
}
