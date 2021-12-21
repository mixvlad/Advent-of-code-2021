package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []rune
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]rune, 0)}
}

func (s *stack) Push(v rune) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Peek() rune {
	l := len(s.s)
	if l == 0 {
		return 0
	}
	res := s.s[l-1]
	return res
}

func (s *stack) Pop() (rune, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Count() int {
	l := len(s.s)

	return l
}

func Contains(what rune, where map[rune]rune) bool {
	for k, _ := range where {
		if k == what {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	costs := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	compCosts := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	maps := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	// mapsInv := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	summ := 0
	completeScores := make([]int, 0, 100)

	for scanner.Scan() {
		s := NewStack()
		curRow := scanner.Text()
		if len(curRow) > 0 {
			corrupted := false
			for _, v := range curRow {
				if Contains(v, maps) {
					s.Push(v)
				} else {
					if s.Count() == 0 || maps[s.Peek()] != v {
						summ += costs[v]
						corrupted = true
						break
					} else {
						s.Pop()
					}
				}

			}

			if !corrupted {
				var compSumm int
				for {
					if s.Count() == 0 {
						break
					}
					x, _ := s.Pop()

					compSumm = (compSumm * 5) + compCosts[maps[x]]
				}
				completeScores = append(completeScores, compSumm)
			}
		}
	}
	fmt.Println(summ)
	sort.Ints(completeScores)
	fmt.Println(completeScores[(len(completeScores)-1)/2])
}
