package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func find(what int, where []int) (idx int) {
	for i, v := range where {
		if v == what {
			return i
		}
	}
	return -1
}

func contains(what string, where string) bool {
	for _, x := range what {
		found := false
		for _, y := range where {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func findStrByLen(lenS int, where []string) (res []string) {
	for _, x := range where {
		if len(x) == lenS {
			res = append(res, x)
		}
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

	knownDigsLen := []int{2, 3, 4, 7}
	count := 0

	for scanner.Scan() {
		curRow := strings.Split(scanner.Text(), " | ")
		if len(curRow) > 1 {
			segments := strings.Split(curRow[1], " ")
			for _, segment := range segments {
				if find(len(segment), knownDigsLen) >= 0 {
					count++
				}
			}
		}

	}

	fmt.Println(count)

	secondPart()
}

func kByV(what int, where map[string]int) string {
	for k, v := range where {
		if v == what {
			return k
		}
	}
	return ""
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStr(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func secondPart() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		curRow := strings.Split(scanner.Text(), " | ")
		if len(curRow) > 1 {
			var wordsDict = map[string]int{}
			input := strings.Split(curRow[0], " ")
			output := strings.Split(curRow[1], " ")

			wordsDict[sortStr(findStrByLen(2, input)[0])] = 1
			wordsDict[sortStr(findStrByLen(4, input)[0])] = 4
			wordsDict[sortStr(findStrByLen(3, input)[0])] = 7
			wordsDict[sortStr(findStrByLen(7, input)[0])] = 8

			// find 9
			str069 := findStrByLen(6, input)
			for k, v := range str069 {
				if contains(kByV(4, wordsDict), v) {
					wordsDict[sortStr(v)] = 9
					str069 = remove(str069, k)
					break
				}
			}

			// find 0
			for k, v := range str069 {
				if contains(kByV(7, wordsDict), v) {
					wordsDict[sortStr(v)] = 0
					str069 = remove(str069, k)
					break
				}
			}

			// find 6
			wordsDict[sortStr(str069[0])] = 6

			// find 3
			str235 := findStrByLen(5, input)
			for k, v := range str235 {
				if contains(kByV(1, wordsDict), v) {
					wordsDict[sortStr(v)] = 3
					str235 = remove(str235, k)
					break
				}
			}

			// find 5
			for k, v := range str235 {
				if contains(v, kByV(9, wordsDict)) {
					wordsDict[sortStr(v)] = 5
					str235 = remove(str235, k)
					break
				}
			}

			// find 2
			wordsDict[sortStr(str235[0])] = 2

			summ := 0
			for k, v := range output {
				x := wordsDict[sortStr(v)]
				y := int(math.Pow(10, float64(3-k)))
				summ += x * y
			}
			count += summ
		}

	}

	fmt.Println(count)
}
