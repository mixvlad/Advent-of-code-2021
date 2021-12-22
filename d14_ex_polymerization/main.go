package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	inputS := scanner.Text()
	rules := map[string]string{}
	input := map[string]int{}

	for scanner.Scan() {
		row := scanner.Text()

		if len(row) > 0 {
			splitedRow := strings.Split(row, " -> ")
			rules[splitedRow[0]] = splitedRow[1]
			input[splitedRow[0]] = 0
		}
	}

	for i := 0; i < len(inputS)-1; i++ {
		input[string(inputS[i:i+2])]++
	}
	countOfSteps := 10
	input = ApplyPoly(input, countOfSteps, rules)

	countS := CountSymb(input)
	countS[rune(inputS[0])]++
	countS[rune(inputS[len(inputS)-1])]++

	fmt.Println("Part one:", Max(countS)-Min(countS))

	countOfSteps = 30
	input = ApplyPoly(input, countOfSteps, rules)

	countS = CountSymb(input)
	countS[rune(inputS[0])]++
	countS[rune(inputS[len(inputS)-1])]++

	fmt.Println("Part two:", Max(countS)-Min(countS))
}

func CountSymb(input map[string]int) (res map[rune]int) {
	res = map[rune]int{}
	for k, v := range input {
		for _, symbK := range k {
			if _, ok := res[symbK]; !ok {
				res[symbK] = v
			} else {
				res[symbK] += v
			}
		}
	}

	for k, v := range res {
		res[k] = v / 2
	}
	return
}

func ApplyPoly(input map[string]int, countOfSteps int, rules map[string]string) map[string]int {
	for step := 0; step < countOfSteps; step++ {
		buf := map[string]int{}

		for k, v := range input {
			if v > 0 {
				newSymb := rules[k]
				if _, ok := buf[string(k[0])+newSymb]; !ok {
					buf[string(k[0])+newSymb] = v
				} else {
					buf[string(k[0])+newSymb] += v
				}
				if _, ok := buf[newSymb+string(k[1])]; !ok {
					buf[newSymb+string(k[1])] = v
				} else {
					buf[newSymb+string(k[1])] += v
				}
			}
		}
		input = buf
	}

	return input
}

func Min(arr map[rune]int) (res int) {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	res = MaxInt
	for _, v := range arr {
		if v < res {
			res = v
		}
	}
	return res
}

func Max(arr map[rune]int) (res int) {
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
