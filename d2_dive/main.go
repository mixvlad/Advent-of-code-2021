package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y := 0, 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) > 0 {
			intVar, _ := strconv.Atoi(words[1])
			direction := words[0]
			if direction == "forward" {
				x += intVar
			}
			if direction == "down" {
				y += intVar
			}
			if direction == "up" {
				y -= intVar
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(x * y)
}
