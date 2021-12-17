package main

import (
	"bufio"
	"container/list"
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
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
	three_measurement()
}

func three_measurement() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	count := 0
	queue := list.New()
	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		queue.PushFront(intVar)
		e := queue.Back()
		if queue.Len() > 3 {
			prevVar, _ := e.Value.(int)
			if intVar > prevVar {
				count++
			}
			queue.Remove(e)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
