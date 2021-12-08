package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var total, previous int
	total = 0
	previous = 0
	file, err := os.Open("Inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if current > previous {
			total += 1
		}
		previous = current
	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
