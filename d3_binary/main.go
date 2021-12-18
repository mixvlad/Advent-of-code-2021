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

	var numArr []int
	for scanner.Scan() {
		curRow := scanner.Text()
		if aLen := len(curRow); numArr == nil && aLen > 0 {
			numArr = make([]int, aLen)
		}
		for i := 0; i < len(curRow); i++ {
			if string(curRow[i]) == "1" {
				numArr[i]++
			} else {
				numArr[i]--
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var res, xorInt uint = 0, 0
	for i := 0; i < len(numArr); i++ {
		if numArr[i] > 0 {
			res += 1 << (len(numArr) - i - 1)
		}
		xorInt += 1 << (len(numArr) - i - 1)
	}

	secondRes := xorInt ^ res

	fmt.Println(secondRes * res) // result
	// fmt.Printf("%b\n", res)
	// fmt.Printf("%b\n", xorInt)
	// fmt.Printf("%b\n", xorInt^res)
	// fmt.Println(0x111111111111)
	secondPart()
}

func secondPart() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	strArr := make([]string, 0, 1000)

	// fill array of rows
	for scanner.Scan() {
		curRow := scanner.Text()
		if len(curRow) > 0 {
			strArr = append(strArr, curRow)
		}
	}

	more := binaryStrToInt(filterList(strArr, 0, true))
	less := binaryStrToInt(filterList(strArr, 0, false))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(more * less) // result
	fmt.Println(more)        // result
	fmt.Println(less)        // result
	// fmt.Printf("%b\n", res)
	// fmt.Printf("%b\n", xorInt)
	// fmt.Printf("%b\n", xorInt^res)
	// fmt.Println(0x111111111111)

}

func binaryStrToInt(str string) (x int64) {
	x, _ = strconv.ParseInt(str, 2, 64)
	return
}

func filterList(list []string, index int, filterWhichMore bool) string {
	if len(list) == 1 || len(list[0]) <= index {
		return list[0]
	}

	zeroes, ones := make([]string, 0, len(list)), make([]string, 0, len(list))
	for _, x := range list {
		if string(x[index]) == "1" {
			ones = append(ones, x)
		} else {
			zeroes = append(zeroes, x)
		}
	}

	if filterWhichMore {
		if len(ones) >= len(zeroes) {
			return filterList(ones, index+1, filterWhichMore)
		}
		return filterList(zeroes, index+1, filterWhichMore)
	} else {
		if len(zeroes) <= len(ones) {
			return filterList(zeroes, index+1, filterWhichMore)
		}
		return filterList(ones, index+1, filterWhichMore)
	}
}
