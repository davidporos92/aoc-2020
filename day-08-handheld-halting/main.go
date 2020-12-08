package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"log"
	"strconv"
	"strings"
)

const (
	SignIncrease = iota
	SignDecrease

	OperationAccumulator = "acc"
	OperationJump        = "jmp"
	OperationNoop        = "nop"
)

type BootCode struct {
	iteration   int
	accumulator int
	nextCommand int
	commands    []*Command
}

type Command struct {
	operation     string
	argumentValue int
	argumentSign  uint
	visited       []int
}

func NewCommand(command string) *Command {
	commandParts := strings.Split(command, " ")
	argumentValue, argumentSign := parseArgument(commandParts[1])

	return &Command{
		operation:     commandParts[0],
		argumentValue: argumentValue,
		argumentSign:  argumentSign,
		visited:       make([]int, 0),
	}
}

func parseArgument(argument string) (int, uint) {
	value, err := strconv.Atoi(strings.Trim(argument, "+-"))
	if err != nil {
		log.Fatalf("Cannot convert %s to int: %v", argument, err)
	}

	if strings.Contains(argument, "-") {
		return value, SignDecrease
	}

	return value, SignIncrease
}

func (bc *BootCode) GetAccBeforeACommandRunTwice() int {
	for {
		actualCommand := bc.commands[bc.nextCommand]
		actualAcc := bc.accumulator
		bc.Exec(bc.commands[bc.nextCommand])
		if len(actualCommand.visited) > 1 {
			return actualAcc
		}

		if bc.nextCommand >= len(bc.commands) {
			break
		}
	}

	return -1
}

func (bc *BootCode) FixCode() bool {
	if !bc.hasInfiniteLoop() {
		return true
	}

	commandList := bc.getPossiblyWrongCommandRunList()
	tryToFixCommandIndex := len(commandList)-1

	for {
		origOp := commandList[tryToFixCommandIndex].operation
		if origOp == OperationNoop {
			commandList[tryToFixCommandIndex].operation = OperationJump
		} else {
			commandList[tryToFixCommandIndex].operation = OperationNoop
		}

		bc.reset()
		if !bc.hasInfiniteLoop() {
			return true
		}

		commandList[tryToFixCommandIndex].operation = origOp

		tryToFixCommandIndex--
		if tryToFixCommandIndex < 0 {
			return false
		}
	}
}

func (bc *BootCode) Exec(c *Command) {
	defer func() {
		c.visited = append(c.visited, bc.iteration)
		bc.iteration++
	}()

	switch c.operation {
	case OperationAccumulator:
		bc.acc(c)
		bc.nextCommand++
	case OperationJump:
		bc.jmp(c)
	case OperationNoop:
		bc.nextCommand++
	}
}

func (bc *BootCode) acc(c *Command) {
	switch c.argumentSign {
	case SignDecrease:
		bc.accumulator -= c.argumentValue
	case SignIncrease:
		bc.accumulator += c.argumentValue
	default:
		log.Fatalf("Unknown argument sign: %v", c.argumentSign)
	}
}

func (bc *BootCode) jmp(c *Command) {
	switch c.argumentSign {
	case SignDecrease:
		bc.nextCommand -= c.argumentValue
	case SignIncrease:
		bc.nextCommand += c.argumentValue
	default:
		log.Fatalf("Unknown argument sign: %v", c.argumentSign)
	}
}

func (bc *BootCode) hasInfiniteLoop() bool {
	acc := bc.GetAccBeforeACommandRunTwice()
	if acc != -1 {
		return true
	}

	return false
}

func (bc *BootCode) reset() {
	bc.resetCommandVisits()
	bc.accumulator = 0
	bc.nextCommand = 0
	bc.iteration = 0
}

func (bc *BootCode) resetCommandVisits() {
	for _, c := range bc.commands {
		c.visited = make([]int, 0)
	}
}

func (bc *BootCode) getPossiblyWrongCommandRunList() []*Command {
	list := make([]*Command, 0, len(bc.commands))

	for _, c := range bc.commands {
		if c.operation != OperationAccumulator && len(c.visited) > 0 {
			list = append(list, c)
		}
	}

	return list
}

func main() {
	bc := &BootCode{
		commands: make([]*Command, 0),
	}

	utils.NewReader("./input-1.dat").MustReadFile(func(command string) {
		bc.commands = append(bc.commands, NewCommand(command))
	})

	fmt.Printf("Solution 1: %d\n", bc.GetAccBeforeACommandRunTwice())

	bc.FixCode()
	fmt.Printf("Solution 2: %d\n", bc.accumulator)
}
