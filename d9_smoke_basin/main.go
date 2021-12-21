package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func isLowest(i, j int, heights [][]int) bool {
	if i > 0 {
		if heights[i-1][j] <= heights[i][j] {
			return false
		}
	}
	if j > 0 {
		if heights[i][j-1] <= heights[i][j] {
			return false
		}
	}
	if i < len(heights)-1 {
		if heights[i+1][j] <= heights[i][j] {
			return false
		}
	}
	if j < len(heights[i])-1 {
		if heights[i][j+1] <= heights[i][j] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heights := make([][]int, 100)
	for i := range heights {
		heights[i] = make([]int, 100)
	}
	var row int
	for scanner.Scan() {

		curRow := scanner.Text()
		if len(curRow) > 0 {
			for k, v := range curRow {
				heights[row][k], _ = strconv.Atoi(string(v))
			}
		}
		row++
	}

	height := 0
	lenBasin := []int{}
	for i := range heights {
		for j := range heights[i] {
			if isLowest(i, j, heights) {
				height += heights[i][j] + 1

				basin := []pair{{i, j}}
				basin = findBasin(heights, basin, i, j)
				lenBasin = append(lenBasin, len(basin))
			}
		}
	}

	fmt.Println(height)
	sort.Slice(lenBasin, func(i, j int) bool {
		return lenBasin[j] < lenBasin[i]
	})
	fmt.Println(lenBasin[0] * lenBasin[1] * lenBasin[2])
}

func Contains(what pair, where []pair) bool {
	for _, x := range where {
		if what == x {
			return true
		}
	}
	return false
}

func Chech(heights [][]int, basin []pair, i, j int) []pair {
	if heights[i][j] < 9 && !Contains(pair{i, j}, basin) {
		basin = append(basin, pair{i, j})
		basin = findBasin(heights, basin, i, j)
	}
	return basin
}

func findBasin(heights [][]int, basin []pair, i, j int) []pair {
	if i > 0 {
		basin = Chech(heights, basin, i-1, j)
	}
	if j > 0 {
		basin = Chech(heights, basin, i, j-1)
	}
	if i < len(heights)-1 {
		basin = Chech(heights, basin, i+1, j)
	}
	if j < len(heights[i])-1 {
		basin = Chech(heights, basin, i, j+1)
	}
	return basin
}

type pair struct {
	x, y int
}
