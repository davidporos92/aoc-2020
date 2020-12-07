package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var bags = make(map[string]*Bag)

type Bag struct {
	name             string
	canContain       map[string]int
	canBeContainedIn []string
}

func NewBag(name string) *Bag {
	bag := &Bag{
		name:             name,
		canContain:       make(map[string]int),
		canBeContainedIn: make([]string, 0),
	}
	bags[name] = bag

	return bag
}

func main() {
	ruleMatcher := regexp.MustCompile(`(\w+\s\w+)\sbags?\scontain\s(.*)`)
	canContainMatcher := regexp.MustCompile(`(\d+)\s(\w+\s\w+)\sbags?`)
	myBagName := "shiny gold"

	utils.NewReader("./input-1.dat").MustReadFile(func(rule string) {
		ruleMatches := ruleMatcher.FindStringSubmatch(rule)
		ruleBag, exist := bags[ruleMatches[1]]
		if !exist {
			ruleBag = NewBag(ruleMatches[1])
		}

		if strings.Contains(ruleMatches[2], "no other bag") {
			return
		}

		ruleBagCanContain := strings.Split(ruleMatches[2], ",")
		for _, canContainRule := range ruleBagCanContain {
			canContainRule = strings.Trim(canContainRule, " ")
			canBeContainedInBagRule := canContainMatcher.FindStringSubmatch(canContainRule)

			qty, err := strconv.Atoi(canBeContainedInBagRule[1])
			if err != nil {
				log.Fatalf("Cannot convert %s to int: %v", canBeContainedInBagRule[1], err)
			}

			ruleBag.canContain[canBeContainedInBagRule[2]] = qty
			canBeContainedInBag, exists := bags[canBeContainedInBagRule[2]]
			if !exists {
				canBeContainedInBag = NewBag(canBeContainedInBagRule[2])
			}
			canBeContainedInBag.canBeContainedIn = append(canBeContainedInBag.canBeContainedIn, ruleBag.name)
		}
	})

	myBag := bags[myBagName]

	fmt.Printf("Solution 1: %d\n", BagsCanHold(myBag.name))
	fmt.Printf("Solution 2: %d\n", NumberOfBagsMustBeContained(myBag))
}

func BagsCanHold(bagName string) int {
	count := 0

	for _, bag := range bags {
		if _, exists := bag.canContain[bagName]; exists {
			count += 1 + BagsCanHold(bag.name)
		}
	}

	return count
}

func NumberOfBagsMustBeContained(bag *Bag) int {
	count := 0

	for containedBag, containedBagCount := range bag.canContain {
		count += containedBagCount + (containedBagCount * NumberOfBagsMustBeContained(bags[containedBag]))
	}

	return count
}
