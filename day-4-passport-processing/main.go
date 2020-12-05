package main

import (
	"fmt"
	"github.com/davidporos92/aoc-2020/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid int
}

func NewPassport(passport string) *Passport {
	p := &Passport{}

	byr := regexp.MustCompile(`byr:([^\s]+)`).FindStringSubmatch(passport)
	if len(byr) > 1 {
		intVal, err := strconv.Atoi(byr[1])
		if err != nil {
			log.Fatalf("Cannot convert byr to int: %v", err)
		}
		p.byr = intVal
	}
	iyr := regexp.MustCompile(`iyr:([^\s]+)`).FindStringSubmatch(passport)
	if len(iyr) > 1 {
		intVal, err := strconv.Atoi(iyr[1])
		if err != nil {
			log.Fatalf("Cannot convert iyr to int: %v", err)
		}
		p.iyr = intVal
	}
	eyr := regexp.MustCompile(`eyr:([^\s]+)`).FindStringSubmatch(passport)
	if len(eyr) > 1 {
		intVal, err := strconv.Atoi(eyr[1])
		if err != nil {
			log.Fatalf("Cannot convert eyr to int: %v", err)
		}
		p.eyr = intVal
	}
	hgt := regexp.MustCompile(`hgt:([^\s]+)`).FindStringSubmatch(passport)
	if len(hgt) > 1 {
		p.hgt = hgt[1]
	}
	hcl := regexp.MustCompile(`hcl:([^\s]+)`).FindStringSubmatch(passport)
	if len(hcl) > 1 {
		p.hcl = hcl[1]
	}
	ecl := regexp.MustCompile(`ecl:([^\s]+)`).FindStringSubmatch(passport)
	if len(ecl) > 1 {
		p.ecl = ecl[1]
	}
	pid := regexp.MustCompile(`pid:([^\s]+)`).FindStringSubmatch(passport)
	if len(pid) > 1 {
		p.pid = pid[1]
	}
	cid := regexp.MustCompile(`cid:([^\s]+)`).FindStringSubmatch(passport)
	if len(cid) > 1 {
		intVal, err := strconv.Atoi(cid[1])
		if err != nil {
			log.Fatalf("Cannot convert cid to int: %v", err)
		}
		p.cid = intVal
	}

	return p
}

func (p *Passport) IsValidSolution1() bool {
	return p.byr != 0 &&
		p.iyr != 0 &&
		p.eyr != 0 &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p *Passport) IsValidSolution2() bool {
	return p.IsValidSolution1() && p.IsByrValid() &&
		p.IsIyrValid() &&
		p.IsEyrValid() &&
		p.IsHgtValid() &&
		p.IsHclValid() &&
		p.IsEclValid() &&
		p.IsPidValid()
}

func (p *Passport) IsByrValid() bool {
	return p.byr >= 1920 && p.byr <= 2002
}

func (p *Passport) IsIyrValid() bool {
	return p.iyr >= 2010 && p.iyr <= 2020
}

func (p *Passport) IsEyrValid() bool {
	return p.eyr >= 2020 && p.eyr <= 2030
}

func (p *Passport) IsHgtValid() bool {
	height, suffix := p.ParseHeight()
	if suffix == "cm" {
		return height >= 150 && height <= 193
	}
	return height >= 59 && height <= 76
}

func (p *Passport) IsHclValid() bool {
	return len(p.hcl) == 7 && regexp.MustCompile(`(?i)#[0-9-a-f]{6}`).MatchString(p.hcl)
}

func (p *Passport) IsEclValid() bool {
	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func (p *Passport) IsPidValid() bool {
	return len(p.pid) == 9 && regexp.MustCompile(`(?i)[0-9]{9}`).MatchString(p.pid)
}

func (p *Passport) ParseHeight() (int, string) {
	var suffix string
	if strings.Contains(p.hgt, "cm") {
		suffix = "cm"
	} else {
		suffix = "in"
	}

	intValue, err := strconv.Atoi(strings.TrimRight(p.hgt, suffix))
	if err != nil {
		log.Fatalf("Cannot convert hgt [%s] to int: %v", p.hgt, err)
	}

	return intValue, suffix
}

func main() {
	var validCountSolution1 int
	var validCountSolution2 int
	passports := make([]*Passport, 0)
	batches := utils.NewReader("./input-1.dat").MustReadStringBatchesFromFile(utils.BatchSeparatorBlankLine)
	for _, batch := range batches {
		passport := NewPassport(batch)
		passports = append(passports, passport)

		if passport.IsValidSolution1() {
			validCountSolution1++
		}

		if passport.IsValidSolution2() {
			validCountSolution2++
		}
	}

	fmt.Printf("Solution 1: %d\n", validCountSolution1)
	fmt.Printf("Solution 2: %d\n", validCountSolution2)
}
