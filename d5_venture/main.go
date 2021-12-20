package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func intersect(l1 line, l2 line, c chan int) {

	c <- 1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]line, 0, 500)

	allLines := make([]line, 0, 500)

	// read all horizontal and vertical lines
	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), " -> ")
		if len(coordinates) > 1 {
			start := strArrToIntArr(strings.Split(coordinates[0], ","))
			end := strArrToIntArr(strings.Split(coordinates[1], ","))
			allLines = append(allLines, line{start[0], start[1], end[0], end[1]})
			if start[0] == end[0] || start[1] == end[1] {
				lines = append(lines, line{start[0], start[1], end[0], end[1]})
			}
		}
	}

	fmt.Println(countDangerousPoints(lines))

	// second part
	fmt.Println(countDangerousPoints(allLines))
}

func countDangerousPoints(lines []line) (summ int) {
	// iterate through lines
	var board [1000][1000]int
	for _, line := range lines {
		dx := line.x2 - line.x1
		dy := line.y2 - line.y1

		if count := Abs(dx); count == Abs(dy) {
			iteratorX := dx / Abs(dx)
			iteratorY := dy / Abs(dy)

			for i := 0; i <= count; i++ {
				board[line.x1+(i*iteratorX)][line.y1+(i*iteratorY)]++
			}
		} else {
			if dx != 0 {
				iteratorX := dx / Abs(dx)
				for i := 0; i <= Abs(dx); i++ {
					board[line.x1+(i*iteratorX)][line.y1]++
				}
			} else {
				iteratorY := dy / Abs(dy)
				for i := 0; i <= Abs(dy); i++ {
					board[line.x1][line.y1+(i*iteratorY)]++
				}
			}
		}
	}

	// summ all
	for i := range board {
		for j := range board[i] {
			if board[i][j] > 1 {
				summ++
			}
		}
	}
	return
}

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

type line struct {
	x1, y1, x2, y2 int
}

func secondPart() {

}
