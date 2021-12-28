package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curRow := scanner.Text()
		fmt.Println(curRow)
	}
	fmt.Println("Part one:", 4+300+3+8000+50+50+700+9000+3+8+500)
	fmt.Println("Part one:", 5+60+50+500+3+8000+30+40+600+9000+8)
	fmt.Println("Part one:", 4+5000+100+700+303+3000+9000+2+500)

	fmt.Println("Part two:")

}
