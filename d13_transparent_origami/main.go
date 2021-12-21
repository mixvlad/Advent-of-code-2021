package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x, y int
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
	dots := map[pair]bool{}
	foldArr := []pair{}

	for scanner.Scan() {
		row := scanner.Text()

		if len(row) > 0 {
			coord := strArrToIntArr(strings.Split(row, ","))
			dots[pair{coord[0], coord[1]}] = true
		} else {
			for scanner.Scan() {
				row = scanner.Text()
				if len(row) > 0 {
					foldS := strings.Split(row, " ")
					foldData := strings.Split(foldS[2], "=")
					value, _ := strconv.Atoi(foldData[1])
					if foldData[0] == "x" {
						foldArr = append(foldArr, pair{value, 0})
					} else {
						foldArr = append(foldArr, pair{0, value})
					}
				}
			}
			break
		}
	}

	for k, v := range foldArr {
		dots = Fold(dots, v)
		if k == 0 {
			fmt.Println("Part one:", len(dots))
		}
	}

	fmt.Println("Part two:")
	for i := 0; i < 10; i++ {
		for j := 0; j < 40; j++ {
			if x, ok := dots[pair{j, i}]; !ok {
				fmt.Print(" ")
			} else {
				if x {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}

}

func Fold(dots map[pair]bool, fold pair) map[pair]bool {
	res := map[pair]bool{}

	for k, _ := range dots {
		if fold.x == 0 {
			if fold.y < k.y {
				newCoord := pair{k.x, fold.y - (k.y - fold.y)}
				res[newCoord] = true
			} else {
				res[k] = true
			}
		} else {
			if fold.x < k.x {
				newCoord := pair{fold.x - (k.x - fold.x), k.y}
				res[newCoord] = true
			} else {
				res[k] = true
			}
		}
	}

	return res
}
