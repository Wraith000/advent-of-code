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
	var horizontal, depth, aim int
	horizontal, depth, aim = 0, 0, 0
	var x []string
	file, err := os.Open("Day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		x = strings.Split(scanner.Text(), " ")
		temp, err := strconv.Atoi(x[1])
		switch x[0] {
		case "forward":
			horizontal += temp
			depth += aim * temp
		case "down":
			//depth += temp
			aim += temp
		case "up":
			//depth -= temp
			aim -= temp
		}

		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}
	fmt.Println(horizontal * depth)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
