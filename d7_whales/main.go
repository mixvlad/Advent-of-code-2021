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

	positions := make([]int, 2000)
	positions2 := make([]int, 2000)

	for scanner.Scan() {
		posArr := strArrToIntArr(strings.Split(scanner.Text(), ","))
		for _, pos := range posArr {
			positions[pos]++
			positions2[pos]++
		}
	}

	fmt.Println(calcFuel(positions))

	//second part
	fmt.Println(calcFuel2(positions2))
}

func calcFuel(positions []int) (fuel int) {
	i := 0
	j := len(positions) - 1
	for {
		if positions[i] <= positions[j] {
			fuel += positions[i]
			positions[i+1] += positions[i]
			i++
		} else {
			fuel += positions[j]
			positions[j-1] += positions[j]
			j--
		}
		if i == j {
			break
		}
	}
	return
}

func cost(crabs []int) (cost int) {
	for k, v := range crabs {
		cost += v * Abs(k-len(crabs)-1)
	}

	return
}

func calcFuel2(positions []int) (fuel int) {
	iCrabs := make([]int, 0, 1000)
	jCrabs := make([]int, 0, 1000)
	i := 0
	j := len(positions) - 1
	for {
		costICrabs := cost(iCrabs)
		costJCrabs := cost(jCrabs)
		if costICrabs+positions[i] <= costJCrabs+positions[j] {
			fuel += costICrabs + positions[i]
			if fuel > 0 {
				iCrabs = append(iCrabs, positions[i])
			}
			i++
		} else {
			fuel += costJCrabs + positions[j]
			if fuel > 0 {
				jCrabs = append(jCrabs, positions[j])
			}

			j--
		}
		if i == j {
			break
		}
	}
	return
}
