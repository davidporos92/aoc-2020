package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"log"
	"regexp"
	"strconv"
)

func main() {
	var solutionOneValidCount int
	var solutionTwoValidCount int

	utils.New("./input-1.dat").MustReadFile(func(line string) {
		parts := regexp.MustCompile(`(\d+)-(\d+)\s([a-zA-Z]):\s([a-zA-Z]+)`).FindStringSubmatch(line)

		// First number is the min occurrence for the first solution
		// but the the character position in the second one
		firstNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(fmt.Sprintf("Cannot convert part 1 to int: %s - %v", parts[1], err))
		}

		// Second number is the max occurrence for the first solution
		// but the the character position in the second one
		secondNumber, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(fmt.Sprintf("Cannot convert part 2 to int: %s - %v", parts[2], err))
		}

		letterMatches := regexp.MustCompile(fmt.Sprintf(`%s`, parts[3])).FindAllString(parts[4], -1)
		if len(letterMatches) >= firstNumber && len(letterMatches) <= secondNumber {
			solutionOneValidCount++
		}

		firstChar := string([]rune(parts[4])[firstNumber-1])
		secondChar := string([]rune(parts[4])[secondNumber-1])
		if (firstChar == parts[3] && secondChar != parts[3]) || (firstChar != parts[3] && secondChar == parts[3]) {
			fmt.Printf("[%s] %s[%d] - %s[%d] (%s)\n", parts[4], string([]rune(parts[4])[firstNumber-1]), firstNumber-1, string([]rune(parts[4])[secondNumber-1]), secondNumber-1, parts[3])
			solutionTwoValidCount++
		}
	})

	fmt.Printf("Solution 1: %d\n", solutionOneValidCount)
	fmt.Printf("Solution 2: %d\n", solutionTwoValidCount)
}
