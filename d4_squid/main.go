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

	// read num order
	scanner.Scan()
	numOrder := strArrToIntArr(strings.Split(scanner.Text(), ","))

	boards := make([][][]bnum, 0, 1000)

	for scanner.Scan() {
		curRow := scanner.Text()
		if len(curRow) > 0 {
			board := make([][]bnum, 5)
			for i := range board {
				board[i] = strOfNumsToArray(curRow)
				scanner.Scan()
				curRow = scanner.Text()
			}
			boards = append(boards, board)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var curNum bnum
	var resBoard *[][]bnum
	for _, curNum = range numOrder {
		for i := range boards {
			if checkBoard(boards[i], curNum.num) {
				resBoard = &boards[i]
				break
			}
		}
		if resBoard != nil {
			break
		}
	}

	fmt.Println(sumBoard(*resBoard) * curNum.num)

	//second_part()
	filledBoards := make([]int, 0, len(boards))
	var lastWinNum int
	for _, curNum = range numOrder {
		for i := range boards {
			if !contains(filledBoards, i) && checkBoard(boards[i], curNum.num) {
				resBoard = &boards[i]
				lastWinNum = curNum.num
				filledBoards = append(filledBoards, i)
			}

		}
		if len(boards) == len(filledBoards) {
			break
		}
	}
	fmt.Println(sumBoard(*resBoard) * lastWinNum)

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sumBoard(board [][]bnum) (sum int) {
	for i := range board {
		for j := range board[i] {
			if !board[i][j].checked {
				sum += board[i][j].num
			}
		}
	}
	return
}

func checkBLine(board [][]bnum, i int, j int, isHorizontal bool) bool {
	for k := range board {
		if (isHorizontal && !board[k][j].checked) || (!isHorizontal && !board[i][k].checked) {
			return false
		}
	}
	return true
}

func isFilled(board [][]bnum, i int, j int) bool {

	return (checkBLine(board, i, j, true)) || (checkBLine(board, i, j, false))
}

func checkBoard(board [][]bnum, num int) bool {

	for i := range board {
		for j := range board[i] {
			if board[i][j].num == num {
				board[i][j].checked = true
				return isFilled(board, i, j)
			}
		}
	}
	return false
}

func strArrToIntArr(strArr []string) (intArr []bnum) {
	intArr = make([]bnum, len(strArr))
	for key, value := range strArr {
		intArr[key].num, _ = strconv.Atoi(value)
	}
	return
}

func strOfNumsToArray(str string) (res []bnum) {
	res = strArrToIntArr(strings.Fields(str))
	return
}

type bnum struct {
	num     int
	checked bool
}

func remove(s [][][]bnum, i int) [][][]bnum {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
