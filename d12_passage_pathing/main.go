package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caves := map[string][]string{}

	for scanner.Scan() {
		curRow := scanner.Text()
		if len(curRow) > 0 {
			connection := strings.Split(scanner.Text(), "-")
			AddConnection(caves, connection[0], connection[1])
			AddConnection(caves, connection[1], connection[0])
		}
	}

	res := unique(GoDeeper(caves, "start", "start"))
	fmt.Println("Part one:", len(res))

	res = unique(GoDeeper2(caves, "start", "START"))
	fmt.Println("Part two:", len(res))
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func GoDeeper(caves map[string][]string, cave string, path string) (res []string) {
	for _, x := range caves[cave] {
		if !strings.Contains(path, x) || IsUpper(x) {
			if x != "end" {
				res = append(res, GoDeeper(caves, x, path+","+x)...)
			} else {
				res = append(res, path+","+x)
			}
		}
	}
	return
}

func GoDeeper2(caves map[string][]string, cave string, path string) (res []string) {
	for _, x := range caves[cave] {

		if path[0] != '!' || !strings.Contains(path, x) || IsUpper(x) {
			if x != "end" {
				if IsLower(x) && strings.Contains(path, x) {
					res = append(res, GoDeeper2(caves, x, "!"+path+","+x)...)
				} else {
					res = append(res, GoDeeper2(caves, x, path+","+x)...)
				}
			} else {
				res = append(res, path+","+x)
			}
		}
	}
	return
}

func AddConnection(caves map[string][]string, x, y string) {
	if y == "start" {
		return
	}
	if _, ok := caves[x]; !ok {
		caves[x] = []string{y}
	} else {
		caves[x] = append(caves[x], y)
	}
}
