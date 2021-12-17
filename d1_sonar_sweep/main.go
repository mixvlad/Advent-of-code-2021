package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	count := 0
	var intPointer *int
	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		if intPointer != nil && intVar > *intPointer {
			count++
		}
		intPointer = &intVar
		// if prevX == nil {

		// }
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
