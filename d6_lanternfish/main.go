package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func strArrToIntArr(strArr []string) (intArr []int) {
	intArr = make([]int, len(strArr))
	for key, value := range strArr {
		intArr[key], _ = strconv.Atoi(value)
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fishList := make([]int, 9)

	// read all horizontal and vertical lines
	for scanner.Scan() {
		fishes := strArrToIntArr(strings.Split(scanner.Text(), ","))
		for _, fish := range fishes {
			fishList[fish]++
		}
	}

	numOfDaysToPass := 80
	for i := 0; i < numOfDaysToPass; i++ {
		fishList = calcFishes(fishList)
	}

	summ := 0
	for _, v := range fishList {
		summ += v
	}
	fmt.Println(summ)

	//second part
	numOfDaysToPass = 256 - numOfDaysToPass
	for i := 0; i < numOfDaysToPass; i++ {
		fishList = calcFishes(fishList)
	}

	summ = 0
	for _, v := range fishList {
		summ += v
	}
	fmt.Println(summ)
}

func calcFishes(fishes []int) (res []int) {
	res = make([]int, len(fishes))

	for i := 0; i < len(fishes); i++ {
		if i == 0 {
			res[8] += fishes[0]
			res[6] += fishes[0]
		} else {
			res[i-1] += fishes[i]
		}
	}
	return
}
