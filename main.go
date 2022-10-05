package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("File.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines[0])
	fmt.Println(lines[(len(lines) - 1)])

	for index, words := range lines {
		if words == "before" {
			value := lines[index+1]
			a, _ := strconv.Atoi(value)
			fmt.Println(lines[a-1])
		}

		if words == "now " {
			b := lines[index-1]
			c := int(b[1]) / len(lines)
			fmt.Println(lines[c-1])
		}

	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())

	file.Close()
	// for _, eachline := range lines {
	// 	fmt.Println(eachline)
	// }
}
